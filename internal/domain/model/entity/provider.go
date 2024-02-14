package entity

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Provider struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `bson:"name"`
	TaskValueKey    string             `bson:"task_value_key"`
	TaskDurationKey string             `bson:"task_duration_key"`
	TaskNameKey     string             `bson:"task_name_key"`
	Url             string             `bson:"url"`
	Tasks           []valueobject.Task `bson:"tasks,omitempty"`
}
