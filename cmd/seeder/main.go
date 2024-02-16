package seeder

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/dto"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/service"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/config"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/env"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/logger"
	mongodb "github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/storage/mongo"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/storage/mongo/query"
	http_client "github.com/bulutcan99/weekly-task-scheduler/internal/transport/http"
	"github.com/goccy/go-json"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	Env *env.ENV
)

func Init() {
	Env = env.ParseEnv()
	logger.Set()
}

func Start() {
	Init()
	http_client.Init()
	slog.Info("Starting seeder...")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	cfg := config.New()
	slog.Info("Config initialized")
	db := mongodb.NewConnection(ctx, cfg.Mongo)
	defer db.Close()
	slog.Info("Database connected!")
	provider := query.NewProviderRepository(db, Env.ProviderCollection)
	task := query.NewTaskRepository(db, Env.TaskCollection)
	slog.Info("Repos initialized")
	providers, err := getProviders(ctx)
	if err != nil {
		slog.Error("Error in getting providers")
		panic(err)
	}
	providerService := service.NewProviderService(provider)
	taskService := service.NewTaskService(task)
	fetcher := http_client.NewFetcher(providerService, taskService)
	for _, provider := range providers {
		err := providerService.AddProvider(ctx, &entity.Provider{
			ID:              provider.ID,
			Name:            provider.Name,
			Url:             provider.Url,
			TaskNameKey:     provider.TaskNameKey,
			TaskValueKey:    provider.TaskValueKey,
			TaskDurationKey: provider.TaskDurationKey,
		})
		if err != nil {
			slog.Error("Error in adding provider")
			panic(err)
		}
	}
	slog.Info("Services initialized")

	err = fetcher.FetchTasksFromMongo(ctx)
	if err != nil {
		slog.Error("Error in scheduler")
		panic(err)
	}
	slog.Info("Seeder finished")
}

func getProviders(ctx context.Context) ([]dto.Provider, error) {
	data, err := os.ReadFile("pkg/mock-response/mock.json")
	if err != nil {
		return nil, err
	}
	var providers []dto.Provider
	err = json.Unmarshal(data, &providers)
	if err != nil {
		return nil, err
	}

	return providers, nil
}
