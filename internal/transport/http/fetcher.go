package http_client

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"github.com/bulutcan99/weekly-task-scheduler/pkg/helper"
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"time"
)

type Fetcher struct {
	providerService interfaces.IProviderService
	taskService     interfaces.ITaskService
}

func NewFetcher(providerService interfaces.IProviderService, taskService interfaces.ITaskService) *Fetcher {
	return &Fetcher{
		providerService: providerService,
		taskService:     taskService,
	}
}

func (s *Fetcher) Fetch(ctx context.Context) error {
	providers, err := s.providerService.GetProviders(ctx)
	if err != nil {
		return err
	}

	for _, provider := range providers {
		_, body, err := SendGetRequest(provider.Url)
		if err != nil {
			continue
		}

		var tasks []map[string]interface{}
		err = json.Unmarshal(body, &tasks)
		if err != nil {
			slog.Error("Error while parsing body: ", err)
			continue
		}

		for _, task := range tasks {
			taskName := task[provider.TaskNameKey].(string)
			difficulty := helper.ConvertToInt(task[provider.TaskValueKey].(float64))
			duration := helper.ConvertToInt(task[provider.TaskDurationKey].(float64))
			taskData := &valueobject.Task{
				ID:         primitive.NewObjectID(),
				ProviderID: provider.ID,
				TaskName:   taskName,
				Difficulty: difficulty,
				Duration:   duration,
				CreatedAt:  time.Now(),
			}
			err = s.taskService.AddTask(ctx, taskData)
		}
	}

	return nil
}
