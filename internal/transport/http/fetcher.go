package http_client

import (
	"context"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"github.com/bulutcan99/weekly-task-scheduler/pkg/helper"
	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"sync"
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

func (s *Fetcher) FetchTasksFromMongo(ctx context.Context) error {
	providers, err := s.providerService.GetProviders(ctx)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(providers))

	for _, provider := range providers {
		wg.Add(1)
		go func(provider *entity.Provider) {
			defer wg.Done()

			_, body, err := SendGetRequest(provider.Url)
			if err != nil {
				errCh <- err
				return
			}

			var tasks []map[string]any
			if err := json.Unmarshal(body, &tasks); err != nil {
				slog.Error("Error while parsing body: ", err)
				errCh <- err
				return
			}

			for _, task := range tasks {
				taskName := task[provider.TaskNameKey].(string)
				difficulty := helper.ConvertToInt(task[provider.TaskValueKey].(float64))
				duration := helper.ConvertToInt(task[provider.TaskDurationKey].(float64))
				taskData := &valueobject.Task{
					ProviderID: provider.ID,
					Name:       taskName,
					Difficulty: difficulty,
					Duration:   duration,
					Intensity:  difficulty * duration,
				}

				if err := s.taskService.AddTask(ctx, taskData); err != nil {
					slog.Error("Error while adding task: ", err)
					errCh <- err
					return
				}
			}
		}(&provider)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Fetcher) FetchTasks(ctx context.Context, provider *entity.Provider) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	wg.Add(1)
	go func(provider *entity.Provider) {
		defer wg.Done()

		_, body, err := SendGetRequest(provider.Url)
		if err != nil {
			errCh <- err
			return
		}

		var tasks []map[string]any
		if err := json.Unmarshal(body, &tasks); err != nil {
			slog.Error("Error while parsing body: ", err)
			errCh <- err
			return
		}

		for _, task := range tasks {
			taskName := task[provider.TaskNameKey].(string)
			difficulty := helper.ConvertToInt(task[provider.TaskValueKey].(float64))
			duration := helper.ConvertToInt(task[provider.TaskDurationKey].(float64))
			taskData := &valueobject.Task{
				ID:         primitive.NewObjectID(),
				ProviderID: provider.ID,
				Name:       taskName,
				Difficulty: difficulty,
				Duration:   duration,
				Intensity:  difficulty * duration,
			}

			if err := s.taskService.AddTask(ctx, taskData); err != nil {
				slog.Error("Error while adding task: ", err)
				errCh <- err
				return
			}
		}
	}(provider)

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
