package service

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type MockProviderRepository struct {
	mock.Mock
}

func (m *MockProviderRepository) InsertProvider(ctx context.Context, provider *entity.Provider) error {
	args := m.Called(ctx, provider)
	return args.Error(0)
}

func (m *MockProviderRepository) UpdateProvider(ctx context.Context, provider *entity.Provider) error {
	args := m.Called(ctx, provider)
	return args.Error(0)
}

func (m *MockProviderRepository) FindProviders(ctx context.Context) ([]entity.Provider, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Provider), args.Error(1)
}

func (m *MockProviderRepository) DeleteProviders(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func (m *MockProviderRepository) GetProviderWithTasks(ctx context.Context, id primitive.ObjectID) (*entity.Provider, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Provider), args.Error(1)
}

func TestProviderService_GetProviders(t *testing.T) {
	mockRepo := new(MockProviderRepository)
	providerService := NewProviderService(mockRepo)
	mockProviders := []entity.Provider{
		{ID: primitive.NewObjectID(), Name: "Provider1"},
		{ID: primitive.NewObjectID(), Name: "Provider2"},
	}
	mockRepo.On("FindProviders", mock.Anything).Return(mockProviders, nil)
	providers, err := providerService.GetProviders(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, mockProviders, providers)
	mockRepo.AssertCalled(t, "FindProviders", mock.Anything)
}

func TestProviderService_AddProvider(t *testing.T) {
	mockRepo := new(MockProviderRepository)
	providerService := NewProviderService(mockRepo)
	mockProvider := &entity.Provider{ID: primitive.NewObjectID(), Name: "NewProvider"}
	mockRepo.On("InsertProvider", mock.Anything, mock.Anything).Return(nil)
	err := providerService.AddProvider(context.TODO(), mockProvider)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "InsertProvider", mock.Anything, mockProvider)
}
