package dto

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Provider struct {
	ID              primitive.ObjectID `json:"id"`
	Name            string             `json:"name"`
	TaskValueKey    string             `json:"task_value_key"`
	TaskDurationKey string             `json:"task_duration_key"`
	TaskNameKey     string             `json:"task_name_key"`
	Url             string             `json:"url"`
	Tasks           []valueobject.Task `json:"tasks,omitempty"`
}
