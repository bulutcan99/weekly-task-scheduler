package mapper

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/dto"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
)

func GetProviderJson(provider []entity.Provider) []dto.Provider {
	var providers []dto.Provider
	for _, provider := range provider {
		providers = append(providers, dto.Provider{
			ID:              provider.ID,
			Name:            provider.Name,
			TaskValueKey:    provider.TaskValueKey,
			TaskDurationKey: provider.TaskDurationKey,
			TaskNameKey:     provider.TaskNameKey,
			Url:             provider.Url,
		})
	}
	return providers
}

func GetTasksJson(tasks []valueobject.Task) []dto.Task {
	var taskJson []dto.Task
	for _, task := range tasks {
		taskJson = append(taskJson, dto.Task{
			ID:         task.ID.Hex(),
			Duration:   task.Duration,
			Difficulty: task.Difficulty,
		})
	}
	return taskJson
}
