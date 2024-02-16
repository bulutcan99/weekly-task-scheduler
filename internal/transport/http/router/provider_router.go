package router

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/transport/http/controller"
	"github.com/gofiber/fiber/v2"
)

func ProviderRoute(r fiber.Router, provider *controller.ProviderController) {
	route := r.Group("/v1")
	route.Get("/providers", provider.GetProviders)
	route.Post("/provider", provider.AddProvider)
	route.Put("/provider", provider.UpdateProvider)
	route.Delete("/provider/:name", provider.DeleteProviders)
	route.Get("/provider/:id", provider.GetProviderWithTasks)
}
