package query

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/config"
	mongodb "github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/storage/mongo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type Mongo struct {
	Host string
	Port int
	Name string
}

func NewMongoConfig(host string, port int, name string) *Mongo {
	return &Mongo{
		Host: host,
		Port: port,
		Name: name,
	}
}

func TestProviderRepository_InsertProvider(t *testing.T) {
	testCases := []struct {
		name          string
		inputProvider *entity.Provider
		expectedErr   error
	}{
		{
			name: "InsertProvider_Success",
			inputProvider: &entity.Provider{
				ID:              primitive.NewObjectID(),
				Name:            "test",
				TaskValueKey:    "random",
				TaskDurationKey: "random",
				TaskNameKey:     "random",
				Url:             "random",
			},
			expectedErr: nil,
		},
		{
			name: "InsertProvider_DuplicateName",
			inputProvider: &entity.Provider{
				ID:              primitive.NewObjectID(),
				Name:            "duplicate",
				TaskValueKey:    "random",
				TaskDurationKey: "random",
				TaskNameKey:     "random",
				Url:             "random",
			},
			expectedErr: nil,
		},
	}

	ctx := context.Background()
	mongoConfig := NewMongoConfig("localhost", 27017, "test")
	mongo := mongodb.NewConnection(ctx, (*config.Mongo)(mongoConfig))
	providerRepo := NewProviderRepository(mongo, "test")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := providerRepo.InsertProvider(ctx, tc.inputProvider)

			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
