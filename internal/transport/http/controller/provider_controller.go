package controller

import (
	"errors"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProviderController struct {
	ProviderService interfaces.IProviderService
}

func NewProviderController(providerService interfaces.IProviderService) *ProviderController {
	return &ProviderController{ProviderService: providerService}
}

type addProviderRequest struct {
	Name            string `json:"name"`
	TaskValueKey    string `json:"task_value_key"`
	TaskDurationKey string `json:"task_duration_key"`
	TaskNameKey     string `json:"task_name_key"`
	Url             string `json:"url"`
}

func (pc *ProviderController) AddProvider(ctx fiber.Ctx) error {
	var request addProviderRequest
	body := ctx.Body()
	err := json.Unmarshal(body, &request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to parse body",
		})
	}

	provider := &entity.Provider{
		Name:            request.Name,
		TaskValueKey:    request.TaskValueKey,
		TaskDurationKey: request.TaskDurationKey,
		TaskNameKey:     request.TaskNameKey,
		Url:             request.Url,
	}

	err = pc.ProviderService.AddProvider(ctx.Context(), provider)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "provider added successfully"})
}

func (pc *ProviderController) GetProviders(ctx fiber.Ctx) error {
	providers, err := pc.ProviderService.GetProviders(ctx.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": providers})
}

type updateProviderRequest struct {
	Name            string `json:"name,omitempty"`
	TaskValueKey    string `json:"task_value_key,omitempty"`
	TaskDurationKey string `json:"task_duration_key,omitempty"`
	TaskNameKey     string `json:"task_name_key,omitempty"`
	Url             string `json:"url,omitempty"`
}

func (pc *ProviderController) UpdateProvider(ctx fiber.Ctx) error {
	var request updateProviderRequest
	body := ctx.Body()
	err := json.Unmarshal(body, &request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "error while trying to parse body",
		})
	}

	provider := &entity.Provider{
		Name:            request.Name,
		TaskValueKey:    request.TaskValueKey,
		TaskDurationKey: request.TaskDurationKey,
		TaskNameKey:     request.TaskNameKey,
		Url:             request.Url,
	}

	err = pc.ProviderService.UpdateProvider(ctx.Context(), provider)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "provider updated successfully"})
}

func (pc *ProviderController) DeleteProviders(ctx fiber.Ctx) error {
	name := ctx.Params("name")
	err := pc.ProviderService.DeleteProviders(ctx.Context(), name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "provider deleted successfully"})
}

func (pc *ProviderController) GetProviderWithTasks(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	provider, err := pc.ProviderService.GetProviderWithTasks(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": provider})
}
