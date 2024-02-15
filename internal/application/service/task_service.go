package service

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskService struct {
	repo repository.ITaskRepository
}

func NewTaskService(repo repository.ITaskRepository) *TaskService {
	return &TaskService{
		repo,
	}
}

func (ts *TaskService) AddTask(ctx context.Context, task *valueobject.Task) error {
	return ts.repo.InsertTask(ctx, task)
}

func (ts *TaskService) GetTasks(ctx context.Context) ([]valueobject.Task, error) {
	opts := options.Find()
	opts.SetSort(bson.M{
		"intensity": -1,
	})

	tasks, err := ts.repo.GetTasks(ctx, opts)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
