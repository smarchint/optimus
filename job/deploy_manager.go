package job

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/odpf/salt/log"
	"github.com/robfig/cron/v3"

	"github.com/odpf/optimus/config"
	"github.com/odpf/optimus/models"
	"github.com/odpf/optimus/store"
	"github.com/odpf/optimus/utils"
)

const (
	deployAssignInterval = "@every 1m"
)

type DeployManager interface {
	Init()
	Deploy(ctx context.Context, projectSpec models.ProjectSpec) (models.DeploymentID, error)
	GetStatus(ctx context.Context, deployID models.DeploymentID) (models.JobDeployment, error)
}

type deployManager struct {
	l      log.Logger
	config config.Deployer

	deployer         Deployer
	deployerCapacity int32
	uuidProvider     utils.UUIDProvider
	deployRepository store.JobDeploymentRepository

	// unbuffered request channel to assign request to deployer
	requestQ chan models.JobDeployment

	assignerScheduler *cron.Cron

	wg sync.WaitGroup
}

func (m *deployManager) Deploy(ctx context.Context, projectSpec models.ProjectSpec) (models.DeploymentID, error) {
	// Check deployment status for the requested Project
	jobDeployment, err := m.deployRepository.GetByStatusAndProjectID(ctx, models.JobDeploymentStatusInQueue, projectSpec.ID)
	if err == nil {
		return jobDeployment.ID, nil
	}

	// Return valid errors
	if !errors.Is(err, store.ErrResourceNotFound) {
		return models.DeploymentID{}, err
	}

	newDeployment, err := m.createNewRequest(ctx, projectSpec)
	return newDeployment.ID, err
}

func (m *deployManager) createNewRequest(ctx context.Context, projectSpec models.ProjectSpec) (models.JobDeployment, error) {
	newDeploymentID, err := m.uuidProvider.NewUUID()
	if err != nil {
		return models.JobDeployment{}, err
	}

	newDeployment := models.JobDeployment{
		ID:      models.DeploymentID(newDeploymentID),
		Project: projectSpec,
		Status:  models.JobDeploymentStatusInQueue,
		Details: models.JobDeploymentDetail{},
	}

	if err := m.deployRepository.Save(ctx, newDeployment); err != nil {
		return models.JobDeployment{}, err
	}
	return newDeployment, nil
}

func (m *deployManager) GetStatus(ctx context.Context, deployID models.DeploymentID) (models.JobDeployment, error) {
	return m.deployRepository.GetByID(ctx, deployID)
}

func (m *deployManager) Init() {
	m.l.Info("starting deployers", "count", m.config.NumWorkers)
	for i := 0; i < m.config.NumWorkers; i++ {
		m.wg.Add(1)
		go m.spawnDeployer(m.deployer)
	}

	// wait until all deployers are ready
	m.wg.Wait()

	if m.assignerScheduler != nil {
		_, err := m.assignerScheduler.AddFunc(deployAssignInterval, m.Assign)
		if err != nil {
			m.l.Fatal("Failed to start deploy assigner", "error", err)
		}
		m.assignerScheduler.Start()
	}
}

// start a deployer goroutine that runs the deployment in background
func (m *deployManager) spawnDeployer(deployer Deployer) {
	// deployer has spawned
	m.wg.Done()
	atomic.AddInt32(&m.deployerCapacity, 1)

	for deployRequest := range m.requestQ {
		atomic.AddInt32(&m.deployerCapacity, -1)

		m.l.Info("deployer picked up the request", "request id", deployRequest)
		ctx, cancelCtx := context.WithTimeout(context.Background(), m.config.WorkerTimeout)
		if err := deployer.Deploy(ctx, deployRequest); err != nil {
			m.l.Error("deployment worker failed to process", "error", err)
		}
		cancelCtx()

		atomic.AddInt32(&m.deployerCapacity, 1)
	}
}

func (m *deployManager) Assign() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), m.config.WorkerTimeout)
	defer cancelCtx()
	m.cancelTimedOutDeployments(ctx)

	if int(atomic.LoadInt32(&m.deployerCapacity)) <= 0 {
		m.l.Debug("deployers are all occupied.")
		return
	}

	m.l.Debug("trying to assign deployment...")
	jobDeployment, err := m.deployRepository.GetFirstExecutableRequest(ctx)
	if err != nil {
		if errors.Is(err, store.ErrResourceNotFound) {
			m.l.Debug("no deployment request found.")
			return
		}
		m.l.Error(fmt.Sprintf("failed to fetch executable deployment request: %s", err.Error()))
		return
	}

	select {
	case m.requestQ <- jobDeployment:
		m.l.Info(fmt.Sprintf("deployer taking request for %s", jobDeployment.ID.UUID()))
		return
	default:
		m.l.Error(fmt.Sprintf("unable to assign deployer to take request %s", jobDeployment.ID.UUID()))
		jobDeployment.Status = models.JobDeploymentStatusCancelled
		if err := m.deployRepository.Update(ctx, jobDeployment); err != nil {
			m.l.Error(fmt.Sprintf("unable to mark job deployment %s as cancelled: %s", jobDeployment.ID.UUID(), err.Error()))
		}
	}
}

func (m *deployManager) cancelTimedOutDeployments(ctx context.Context) {
	inProgressDeployments, err := m.deployRepository.GetByStatus(ctx, models.JobDeploymentStatusInProgress)
	if err != nil {
		m.l.Error(fmt.Sprintf("failed to fetch in progress deployments: %s", err.Error()))
		return
	}

	for _, deployment := range inProgressDeployments {
		// check state / timed out deployment, mark as cancelled
		if time.Since(deployment.UpdatedAt).Minutes() > m.config.WorkerTimeout.Minutes() {
			deployment.Status = models.JobDeploymentStatusCancelled
			if err := m.deployRepository.Update(ctx, deployment); err != nil {
				m.l.Error(fmt.Sprintf("failed to mark timed out deployment as cancelled: %s", err.Error()))
			}
		}
	}
}

// NewDeployManager constructs a new instance of Job Deployment Manager
func NewDeployManager(l log.Logger, deployerConfig config.Deployer, deployer Deployer, uuidProvider utils.UUIDProvider,
	deployRepository store.JobDeploymentRepository, assignerScheculer *cron.Cron) *deployManager {
	mgr := &deployManager{
		requestQ:          make(chan models.JobDeployment),
		l:                 l,
		config:            deployerConfig,
		deployerCapacity:  0,
		deployer:          deployer,
		uuidProvider:      uuidProvider,
		deployRepository:  deployRepository,
		assignerScheduler: assignerScheculer,
	}
	mgr.Init()
	return mgr
}
