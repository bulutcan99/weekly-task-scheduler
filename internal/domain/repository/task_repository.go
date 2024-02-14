package repository

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ITaskRepository interface {
	InsertTask(ctx context.Context, task *valueobject.Task) error
	GetTasks(ctx context.Context, opt *options.FindOptions) ([]valueobject.Task, error)
}
