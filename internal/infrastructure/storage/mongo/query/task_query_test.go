package query

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/config"
	mongodb "github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/storage/mongo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestTaskRepository_InsertTask(t *testing.T) {
	testCases := []struct {
		name        string
		inputTask   *valueobject.Task
		expectedErr error
	}{
		{
			name: "InsertTask_Success",
			inputTask: &valueobject.Task{
				ID:         primitive.NewObjectID(),
				ProviderID: primitive.NewObjectID(),
				Name:       "ID-1",
				Difficulty: 3,
				Duration:   2,
				Intensity:  6,
			},
			expectedErr: nil,
		},
	}

	ctx := context.Background()
	mongoConfig := NewMongoConfig("localhost", 27017, "test")
	mongo := mongodb.NewConnection(ctx, (*config.Mongo)(mongoConfig))
	taskRepo := NewTaskRepository(mongo, "test")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := taskRepo.CreateOrUpdate(ctx, tc.inputTask, nil)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
