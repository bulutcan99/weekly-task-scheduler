package http_client

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type MockProviderService struct {
	mock.Mock
}

func (m *MockProviderService) GetProviders(ctx context.Context) ([]entity.Provider, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Provider), args.Error(1)
}

func (m *MockProviderService) AddProvider(ctx context.Context, provider *entity.Provider) error {
	args := m.Called(ctx, provider)
	return args.Error(0)
}

func (m *MockProviderService) UpdateProvider(ctx context.Context, provider *entity.Provider) error {
	args := m.Called(ctx, provider)
	return args.Error(0)
}

func (m *MockProviderService) DeleteProviders(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func (m *MockProviderService) GetProviderWithTasks(ctx context.Context, id string) (*entity.Provider, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Provider), args.Error(1)
}

// MockTaskService is a mock implementation for ITaskService.
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetTasks(ctx context.Context) ([]valueobject.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]valueobject.Task), args.Error(1)
}

func (m *MockTaskService) AddTask(ctx context.Context, task *valueobject.Task) error {
	args := m.Called(ctx, task)
	return args.Error(0)
}

func TestFetcher_FetchTasks(t *testing.T) {
	tests := []struct {
		name       string
		mockFields func(providerService *MockProviderService, taskService *MockTaskService)
		wantErr    bool
	}{
		{
			name: "Test FetchTasks",
			mockFields: func(providerService *MockProviderService, taskService *MockTaskService) {
				provider := &entity.Provider{ID: primitive.NewObjectID(), Url: "random", TaskNameKey: "random", TaskValueKey: "random", TaskDurationKey: "random"}
				providerService.On("GetProviders", mock.Anything).Return([]entity.Provider{*provider}, nil)
				taskService.On("AddTask", mock.Anything, mock.AnythingOfType("*valueobject.Task")).Return(nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			providerService := new(MockProviderService)
			taskService := new(MockTaskService)

			tt.mockFields(providerService, taskService)

			s := &Fetcher{
				providerService: providerService,
				taskService:     taskService,
			}

			err := s.FetchTasks(context.Background(), &entity.Provider{})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			providerService.AssertExpectations(t)
			taskService.AssertExpectations(t)
		})
	}
}
func TestFetcher_FetchTasksFromMongo(t *testing.T) {
	type fields struct {
		providerService interfaces.IProviderService
		taskService     interfaces.ITaskService
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(providerService *MockProviderService, taskService *MockTaskService)
		wantErr bool
	}{
		{
			name: "Test FetchTasksFromMongo",
			fields: fields{
				providerService: new(MockProviderService),
				taskService:     new(MockTaskService),
			},
			args:    args{ctx: context.TODO()},
			mock:    func(providerService *MockProviderService, taskService *MockTaskService) {},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Fetcher{
				providerService: tt.fields.providerService,
				taskService:     tt.fields.taskService,
			}
			if tt.mock != nil {
				tt.mock(tt.fields.providerService.(*MockProviderService), tt.fields.taskService.(*MockTaskService))
			}
			if err := s.FetchTasksFromMongo(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("FetchTasksFromMongo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
