package repository

import (
	"context"

	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IProviderRepository interface {
	InsertProvider(ctx context.Context, provider *entity.Provider) error
	UpdateProvider(ctx context.Context, provider *entity.Provider) error
	FindProviders(ctx context.Context) ([]entity.Provider, error)
	DeleteProviders(ctx context.Context, name string) error
	GetProvider(ctx context.Context, filter any) (*entity.Provider, error)
	GetProviderWithTasks(ctx context.Context, id primitive.ObjectID) (*entity.Provider, error)
}
