package aggregate

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
)

type DeveloperTask struct {
	Developer *entity.Developer
	Tasks     *[]valueobject.Task
}
