package aggregate

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
)

type DeveloperTask struct {
	Name   string
	Tasks  []valueobject.Task
	Weeks  int
	Total  int
	Remain int
}
