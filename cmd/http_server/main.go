package http_server

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/service"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/config"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/env"
	fiber_go "github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/fiber"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/logger"
	mongodb "github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/storage/mongo"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/storage/mongo/query"
	http_client "github.com/bulutcan99/weekly-task-scheduler/internal/transport/http"
	"github.com/bulutcan99/weekly-task-scheduler/internal/transport/http/controller"
	"github.com/bulutcan99/weekly-task-scheduler/internal/transport/http/router"
	"github.com/gofiber/fiber/v3"
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
	slog.Info("Starting server...")
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

	providerService := service.NewProviderService(provider)
	taskService := service.NewTaskService(task)
	fetcher := http_client.NewFetcher(providerService, taskService)
	slog.Info("Services initialized")

	cfgFiber := fiber_go.ConfigFiber()
	app := fiber.New(cfgFiber)
	slog.Info("Fiber initialized")
	router.ProviderRoute(app, controller.NewProviderController(providerService, taskService, fetcher))
	router.TaskRoute(app, controller.NewTaskController(taskService))
	slog.Info("Routers initialized")
	fiber_go.FiberListen(ctx, app)
}
