package store

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/odpf/optimus/models"
)

var (
	ErrResourceNotFound = errors.New("resource not found")
)

// JobSpecRepository represents a storage interface for Job specifications
type JobSpecRepository interface {
	Save(models.JobSpec) error
	GetByName(string) (models.JobSpec, error)
	GetAll() ([]models.JobSpec, error)
	Delete(string) error
	GetByDestination(string) (models.JobSpec, models.ProjectSpec, error)
}

// ProjectRepository represents a storage interface for registered projects
type ProjectRepository interface {
	Save(models.ProjectSpec) error
	GetByName(string) (models.ProjectSpec, error)
	GetAll() ([]models.ProjectSpec, error)
}

// ProjectSecretRepository stores secrets attached to projects
type ProjectSecretRepository interface {
	Save(item models.ProjectSecretItem) error
	GetByName(string) (models.ProjectSecretItem, error)
	GetAll() ([]models.ProjectSecretItem, error)
}

// JobRepository represents a storage interface for compiled specifications for
// JobSpecs
type JobRepository interface {
	Save(context.Context, models.Job) error
	GetByName(context.Context, string) (models.Job, error)
	GetAll(context.Context) ([]models.Job, error)
	ListNames(context.Context) ([]string, error)
	Delete(context.Context, string) error
}

// InstanceSpecRepository represents a storage interface for Job runs generated by
// a running instance of job
type InstanceSpecRepository interface {
	Save(models.InstanceSpec) error
	GetByScheduledAt(time.Time) (models.InstanceSpec, error)

	// Clear will not delete the record but will reset all the run details
	Clear(time.Time) error
}

// ResourceSpecRepository represents a storage interface for Respource specifications
type ResourceSpecRepository interface {
	Save(models.ResourceSpec) error
	GetByName(string) (models.ResourceSpec, error)
	GetAll() ([]models.ResourceSpec, error)
	Delete(string) error
}

// ObjectWriter can be used to write in s3 compatible storage interfaces like
// aws s3, gcs, digitalocean buckets, etc.
type ObjectWriter interface {
	NewWriter(ctx context.Context, bucket, path string) (io.WriteCloser, error)
}

// ObjectReader similar to objectWriter but for reading
type ObjectReader interface {
	NewReader(bucket, path string) (io.ReadCloser, error)
}
