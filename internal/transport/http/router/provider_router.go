package router

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/transport/http/controller"
	"github.com/gofiber/fiber/v3"
)

func ProviderRoute(r fiber.Router, provider *controller.ProviderController) {
	route := r.Group("/v1/provider")
	route.Get("/", provider.GetProviders)
	route.Post("/add", provider.AddProvider)
	route.Put("/update", provider.UpdateProvider)
	route.Delete("/delete/:name", provider.DeleteProviders)
	route.Get("/:id", provider.GetProviderWithTasks)
}
