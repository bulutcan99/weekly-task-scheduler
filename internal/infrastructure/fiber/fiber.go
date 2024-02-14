package fiber_go

import (
	"context"
	"fmt"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/env"
	"github.com/gofiber/fiber/v3"
	"log"
	"log/slog"
)

var (
	Host = &env.Env.Host
	Port = &env.Env.ServerPort
)

func FiberListen(ctx context.Context, a *fiber.App) {
	idleConnsClosed := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}
		slog.Info("Server is shutting down...")
		close(idleConnsClosed)
	}()

	fiberConnURL := fmt.Sprintf("%s:%d", *Host, *Port)
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
