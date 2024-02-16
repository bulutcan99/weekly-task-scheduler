package controller

import (
	"errors"
	"fmt"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/dto"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/mapper"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	http_client "github.com/bulutcan99/weekly-task-scheduler/internal/transport/http"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProviderController struct {
	ProviderService interfaces.IProviderService
	TaskService     interfaces.ITaskService
	HttpClient      *http_client.Fetcher
}

func NewProviderController(providerService interfaces.IProviderService, taskService interfaces.ITaskService, fetcher *http_client.Fetcher) *ProviderController {
	return &ProviderController{
		ProviderService: providerService,
		TaskService:     taskService,
		HttpClient:      fetcher,
	}
}

// @Summary Add a new provider
// @Description Add a new provider to add tasks
// @ID insert-provider
// @Accept json
// @Produce json
// @Param request body dto.AddProviderRequest true "Add provider request"
// @Success 201 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/provider [post]
func (pc *ProviderController) AddProvider(ctx *fiber.Ctx) error {
	var request dto.AddProviderRequest
	body := ctx.Body()
	err := json.Unmarshal(body, &request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: true, Msg: "Error while trying to parse body"})
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: true, Msg: err.Error()})

	}

	err = pc.HttpClient.FetchTasks(ctx.Context(), provider)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: true, Msg: err.Error()})

	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.SuccessResponse{Message: "Provider added successfully"})
}

// @Summary Get all the providers
// @Description Get all the providers from the database
// @ID get-providers
// @Produce json
// @Success 200 {object} []dto.Provider
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/providers [get]
func (pc *ProviderController) GetProviders(ctx *fiber.Ctx) error {
	providers, err := pc.ProviderService.GetProviders(ctx.Context())
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: true, Msg: err.Error()})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: true, Msg: err.Error()})
	}
	providerJson := mapper.GetProviderJson(providers)
	fmt.Println(providerJson)
	return ctx.Status(fiber.StatusOK).JSON(providerJson)
}

type updateProviderRequest struct {
	Name            string `json:"name,omitempty"`
	TaskValueKey    string `json:"task_value_key,omitempty"`
	TaskDurationKey string `json:"task_duration_key,omitempty"`
	TaskNameKey     string `json:"task_name_key,omitempty"`
	Url             string `json:"url,omitempty"`
}

func (pc *ProviderController) UpdateProvider(ctx *fiber.Ctx) error {
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

func (pc *ProviderController) DeleteProviders(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	err := pc.ProviderService.DeleteProviders(ctx.Context(), name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "provider deleted successfully"})
}

func (pc *ProviderController) GetProviderWithTasks(ctx *fiber.Ctx) error {
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
