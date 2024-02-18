package service

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) InsertTask(ctx context.Context, task valueobject.Task) error {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockTaskRepository) GetTasks(ctx context.Context, opt *options.FindOptions) ([]valueobject.Task, error) {
	args := m.Called(ctx, opt)
	return args.Get(0).([]valueobject.Task), args.Error(1)
}

func TestTaskService_AddTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskService := NewTaskService(mockRepo)
	mockRepo.On("CreateOrUpdate", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.SingleResult{}, nil)

	task := &valueobject.Task{
		ProviderID: primitive.NewObjectID(),
		Name:       "TaskName",
		Difficulty: 3,
		Duration:   5,
		Intensity:  15,
	}

	err := taskService.AddTask(context.TODO(), task)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "CreateOrUpdate", mock.Anything, bson.M{"provider_id": task.ProviderID, "name": task.Name}, task)
}

func TestTaskService_GetTasks(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskService := NewTaskService(mockRepo)
	mockTasks := []valueobject.Task{
		{ID: primitive.NewObjectID(), Name: "Task1", Intensity: 10},
		{ID: primitive.NewObjectID(), Name: "Task2", Intensity: 20},
	}
	mockRepo.On("GetTasks", mock.Anything, mock.Anything).Return(mockTasks, nil)
	tasks, err := taskService.GetTasks(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, mockTasks, tasks)
	mockRepo.AssertCalled(t, "GetTasks", mock.Anything, mock.AnythingOfType("*options.FindOptions"))
}
