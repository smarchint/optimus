package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/hashicorp/go-hclog"
	hPlugin "github.com/hashicorp/go-plugin"
	"github.com/odpf/salt/log"
	"github.com/prometheus/client_golang/prometheus"
	slackapi "github.com/slack-go/slack"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	v1handler "github.com/odpf/optimus/api/handler/v1beta1"
	pb "github.com/odpf/optimus/api/proto/odpf/optimus/core/v1beta1"
	jobRunCompiler "github.com/odpf/optimus/compiler"
	"github.com/odpf/optimus/config"
	"github.com/odpf/optimus/datastore"
	"github.com/odpf/optimus/ext/notify/slack"
	"github.com/odpf/optimus/job"
	"github.com/odpf/optimus/models"
	"github.com/odpf/optimus/plugin"
	"github.com/odpf/optimus/service"
	"github.com/odpf/optimus/store/postgres"
	"github.com/odpf/optimus/utils"
)

type setupFn func() error

type OptimusServer struct {
	conf   config.ServerConfig
	logger log.Logger

	appKey models.ApplicationKey
	dbConn *gorm.DB

	serverAddr string
	grpcServer *grpc.Server
	httpServer *http.Server

	cleanupFn []func()
}

func New(conf config.ServerConfig) (*OptimusServer, error) {
	addr := fmt.Sprintf("%s:%d", conf.Serve.Host, conf.Serve.Port)
	server := &OptimusServer{
		conf:       conf,
		serverAddr: addr,
	}

	if err := checkRequiredConfigs(conf.Serve); err != nil {
		return server, err
	}

	setupFns := []setupFn{
		server.setupLogger,
		server.setupPlugins,
		server.setupTelemetry,
		server.setupAppKey,
		server.setupDB,
		server.setupGRPCServer,
		server.setupHandlers,
		server.setupMonitoring,
		server.setupHTTPProxy,
	}

	for _, fn := range setupFns {
		if err := fn(); err != nil {
			return server, err
		}
	}

	server.logger.Info("Starting Optimus", "version", config.BuildVersion)
	server.startListening()

	return server, nil
}

func (s *OptimusServer) setupLogger() error {
	s.logger = log.NewLogrus(
		log.LogrusWithLevel(s.conf.Log.Level.String()),
		log.LogrusWithWriter(os.Stderr),
	)
	return nil
}

func (s *OptimusServer) setupPlugins() error {
	pluginLogLevel := hclog.Info
	if s.conf.Log.Level == config.LogLevelDebug {
		pluginLogLevel = hclog.Debug
	}

	pluginLoggerOpt := &hclog.LoggerOptions{
		Name:   "optimus",
		Output: os.Stdout,
		Level:  pluginLogLevel,
	}
	pluginLogger := hclog.New(pluginLoggerOpt)
	s.cleanupFn = append(s.cleanupFn, hPlugin.CleanupClients)

	// discover and load plugins.
	return plugin.Initialize(pluginLogger)
}

func (s *OptimusServer) setupTelemetry() error {
	teleShutdown, err := config.InitTelemetry(s.logger, s.conf.Telemetry)
	if err != nil {
		return err
	}

	s.cleanupFn = append(s.cleanupFn, teleShutdown)
	return nil
}

func (s *OptimusServer) setupAppKey() error {
	var err error
	s.appKey, err = models.NewApplicationSecret(s.conf.Serve.AppKey)
	if err != nil {
		return fmt.Errorf("NewApplicationSecret: %w", err)
	}
	return nil
}

func (s *OptimusServer) setupDB() error {
	var err error
	if err := postgres.Migrate(s.conf.Serve.DB.DSN); err != nil {
		return fmt.Errorf("postgres.Migrate: %w", err)
	}
	// TODO: Connect should accept DBConfig
	s.dbConn, err = postgres.Connect(s.conf.Serve.DB.DSN, s.conf.Serve.DB.MaxIdleConnection,
		s.conf.Serve.DB.MaxOpenConnection, s.logger.Writer())
	if err != nil {
		return fmt.Errorf("postgres.Connect: %w", err)
	}

	return nil
}

func (s *OptimusServer) setupGRPCServer() error {
	var err error
	s.grpcServer, err = setupGRPCServer(s.logger)
	return err
}

func (s *OptimusServer) setupMonitoring() error {
	grpc_prometheus.Register(s.grpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram(grpc_prometheus.WithHistogramBuckets(prometheus.DefBuckets))
	return nil
}

func (s *OptimusServer) setupHTTPProxy() error {
	srv, cleanup, err := prepareHTTPProxy(s.serverAddr, s.grpcServer)
	s.httpServer = srv
	s.cleanupFn = append(s.cleanupFn, cleanup)
	return err
}

func (s *OptimusServer) startListening() {
	// run our server in a goroutine so that it doesn't block to wait for termination requests
	go func() {
		s.logger.Info("Listening at", "address", s.serverAddr)
		if err := s.httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				s.logger.Fatal("server error", "error", err)
			}
		}
	}()
}

func (s *OptimusServer) Shutdown() {
	s.logger.Info("Shutting down server")
	if s.httpServer != nil {
		// Create a deadline to wait for server
		ctxProxy, cancelProxy := context.WithTimeout(context.Background(), shutdownWait)
		defer cancelProxy()

		if err := s.httpServer.Shutdown(ctxProxy); err != nil {
			s.logger.Error("Error in proxy shutdown", err)
		}
	}

	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	for _, fn := range s.cleanupFn {
		fn() // Todo: log all the errors from cleanup before exit
	}

	if s.dbConn != nil {
		sqlConn, err := s.dbConn.DB()
		if err != nil {
			s.logger.Error("Error while getting sqlConn", err)
		} else if err := sqlConn.Close(); err != nil {
			s.logger.Error("Error in sqlConn.Close", err)
		}
	}

	s.logger.Info("Server shutdown complete")
}

func (s *OptimusServer) setupHandlers() error {
	projectRepoFac := &projectRepoFactory{
		db:   s.dbConn,
		hash: s.appKey,
	}

	projectSecretRepo := postgres.NewSecretRepository(s.dbConn, s.appKey)
	namespaceSpecRepoFac := &namespaceRepoFactory{
		db:   s.dbConn,
		hash: s.appKey,
	}
	projectJobSpecRepoFac := &projectJobSpecRepoFactory{
		db: s.dbConn,
	}

	scheduler, err := initScheduler(s.conf)
	if err != nil {
		return err
	}
	models.BatchScheduler = scheduler // TODO: remove global

	engine := jobRunCompiler.NewGoEngine()
	// services
	projectService := service.NewProjectService(projectRepoFac)
	namespaceService := service.NewNamespaceService(projectService, namespaceSpecRepoFac)
	secretService := service.NewSecretService(projectService, namespaceService, projectSecretRepo)
	pluginService := service.NewPluginService(secretService, models.PluginRegistry, engine, s.logger)

	// registered job store repository factory
	jobSpecRepoFac := jobSpecRepoFactory{
		db:                    s.dbConn,
		projectJobSpecRepoFac: *projectJobSpecRepoFac,
	}

	jobDependencyRepo := postgres.NewJobDependencyRepository(s.dbConn)
	dependencyResolver := job.NewDependencyResolver(projectJobSpecRepoFac, jobDependencyRepo, pluginService)
	priorityResolver := job.NewPriorityResolver()

	replaySpecRepoFac := &replaySpecRepoRepository{
		db:             s.dbConn,
		jobSpecRepoFac: jobSpecRepoFac,
	}
	replayWorkerFactory := &replayWorkerFact{
		replaySpecRepoFac: replaySpecRepoFac,
		scheduler:         scheduler,
		logger:            s.logger,
	}
	replayValidator := job.NewReplayValidator(scheduler)
	replaySyncer := job.NewReplaySyncer(
		s.logger,
		replaySpecRepoFac,
		projectRepoFac,
		scheduler,
		func() time.Time {
			return time.Now().UTC()
		},
	)

	replayManager := job.NewManager(s.logger, replayWorkerFactory, replaySpecRepoFac, utils.NewUUIDProvider(), job.ReplayManagerConfig{
		NumWorkers:    s.conf.Serve.Replay.NumWorkers,
		WorkerTimeout: s.conf.Serve.Replay.WorkerTimeout,
		RunTimeout:    s.conf.Serve.Replay.RunTimeout,
	}, scheduler, replayValidator, replaySyncer)

	notificationContext, cancelNotifiers := context.WithCancel(context.Background())
	s.cleanupFn = append(s.cleanupFn, cancelNotifiers)
	eventService := job.NewEventService(s.logger, map[string]models.Notifier{
		"slack": slack.NewNotifier(notificationContext, slackapi.APIURL,
			slack.DefaultEventBatchInterval,
			func(err error) {
				s.logger.Error("slack error accumulator", "error", err)
			},
		),
	})

	deployer := job.NewDeployer(dependencyResolver, priorityResolver, scheduler, namespaceService)

	// runtime service instance over grpc
	manualScheduler := models.ManualScheduler
	jobService := job.NewService(
		&jobSpecRepoFac,
		scheduler,
		manualScheduler,
		jobSpecAssetDump(engine),
		dependencyResolver,
		priorityResolver,
		projectJobSpecRepoFac,
		replayManager,
		namespaceService,
		projectService,
		deployer,
		pluginService,
	)

	jobrunRepoFac := &jobRunRepoFactory{
		db: s.dbConn,
	}

	// job run service
	jobRunService := service.NewJobRunService(
		jobrunRepoFac,
		func() time.Time {
			return time.Now().UTC()
		},
		models.BatchScheduler,
		pluginService,
	)

	progressObs := &pipelineLogObserver{
		log: s.logger,
	}

	projectResourceSpecRepoFac := projectResourceSpecRepoFactory{
		db: s.dbConn,
	}
	resourceSpecRepoFac := resourceSpecRepoFactory{
		db:                         s.dbConn,
		projectResourceSpecRepoFac: projectResourceSpecRepoFac,
	}
	backupRepoFac := backupRepoFactory{
		db: s.dbConn,
	}
	dataStoreService := datastore.NewService(&resourceSpecRepoFac, &projectResourceSpecRepoFac, models.DatastoreRegistry, utils.NewUUIDProvider(), &backupRepoFac, pluginService)
	// adapter service
	adapterService := v1handler.NewAdapter(models.PluginRegistry, models.DatastoreRegistry)

	jobConfigCompiler := jobRunCompiler.NewJobConfigCompiler(engine)
	assetCompiler := jobRunCompiler.NewJobAssetsCompiler(engine, models.PluginRegistry)
	runInputCompiler := jobRunCompiler.NewJobRunInputCompiler(jobConfigCompiler, assetCompiler)

	// secret service
	pb.RegisterSecretServiceServer(s.grpcServer, v1handler.NewSecretServiceServer(s.logger, secretService))
	// resource service
	pb.RegisterResourceServiceServer(s.grpcServer, v1handler.NewResourceServiceServer(s.logger,
		dataStoreService,
		namespaceService,
		adapterService,
		progressObs))
	// replay service
	pb.RegisterReplayServiceServer(s.grpcServer, v1handler.NewReplayServiceServer(s.logger,
		jobService,
		namespaceService,
		adapterService,
		projectService))
	// project service
	pb.RegisterProjectServiceServer(s.grpcServer, v1handler.NewProjectServiceServer(s.logger,
		adapterService,
		projectService))
	// namespace service
	pb.RegisterNamespaceServiceServer(s.grpcServer, v1handler.NewNamespaceServiceServer(s.logger,
		adapterService,
		namespaceService))
	// job Spec service
	pb.RegisterJobSpecificationServiceServer(s.grpcServer, v1handler.NewJobSpecServiceServer(s.logger,
		jobService,
		adapterService,
		projectService,
		namespaceService,
		progressObs))
	// job run service
	pb.RegisterJobRunServiceServer(s.grpcServer, v1handler.NewJobRunServiceServer(s.logger,
		jobService,
		projectService,
		namespaceService,
		secretService,
		adapterService,
		jobRunService,
		runInputCompiler,
		models.BatchScheduler))
	// backup service
	pb.RegisterBackupServiceServer(s.grpcServer, v1handler.NewBackupServiceServer(s.logger,
		jobService,
		dataStoreService,
		namespaceService,
		projectService))
	// runtime service instance over grpc
	pb.RegisterRuntimeServiceServer(s.grpcServer, v1handler.NewRuntimeServiceServer(
		s.logger,
		config.BuildVersion,
		jobService,
		eventService,
		namespaceService,
	))

	cleanupCluster, err := initPrimeCluster(s.logger, s.conf, jobrunRepoFac, s.dbConn)
	if err != nil {
		return err
	}

	s.cleanupFn = append(s.cleanupFn, func() {
		replayManager.Close() // err is nil
	})
	s.cleanupFn = append(s.cleanupFn, cleanupCluster)
	s.cleanupFn = append(s.cleanupFn, func() {
		err = eventService.Close()
		if err != nil {
			s.logger.Error("Error while closing event service: %s", err)
		}
	})

	return nil
}
