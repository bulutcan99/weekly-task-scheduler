package interfaces

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
)

type ITaskService interface {
	AddTask(ctx context.Context, task *valueobject.Task) error
	GetTasks(ctx context.Context) ([]valueobject.Task, error)
}
