package repository

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ITaskRepository interface {
	CreateOrUpdate(ctx context.Context, filter any, update any) *mongo.SingleResult
	GetTasks(ctx context.Context, opt *options.FindOptions) ([]valueobject.Task, error)
}
