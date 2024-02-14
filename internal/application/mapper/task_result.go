package mapper

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/dto"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
)

func TaskJsonProducer(alltasks []valueobject.Task) []dto.TaskJson {
	tasks := make([]dto.TaskJson, 0)

	for _, task := range alltasks {
		taskJson := dto.TaskJson{
			ID:         task.ID.Hex(),
			Difficulty: task.Difficulty,
			Duration:   task.Duration,
		}
		tasks = append(tasks, taskJson)
	}

	return tasks
}
