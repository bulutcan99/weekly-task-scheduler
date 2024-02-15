package aggregate

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
)

type DeveloperTask struct {
	DeveloperName string
	Tasks         *[]valueobject.Task
	Weeks         int
}
