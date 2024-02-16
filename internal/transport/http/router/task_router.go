package router

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/transport/http/controller"
	"github.com/gofiber/fiber/v2"
)

func TaskRoute(r fiber.Router, task *controller.TaskController) {
	route := r.Group("/v1")
	route.Get("/tasks", task.GetTasks)
	route.Post("/assigne-task", task.AssignTask)
}
