package interfaces

import (
	"context"

	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
)

type IProviderService interface {
	AddProvider(ctx context.Context, provider *entity.Provider) error
	UpdateProvider(ctx context.Context, provider *entity.Provider) error
	GetProviders(ctx context.Context) ([]entity.Provider, error)
	DeleteProviders(ctx context.Context, name string) error
	GetProviderWithTasks(ctx context.Context, id string) (*entity.Provider, error)
}
