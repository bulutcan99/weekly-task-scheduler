package entity

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Provider struct {
	ID              primitive.ObjectID `bson:"_id"`
	Name            string             `bson:"name,required"`
	TaskValueKey    string             `bson:"task_value_key,required"`
	TaskDurationKey string             `bson:"task_duration_key,required"`
	TaskNameKey     string             `bson:"task_name_key,required"`
	Url             string             `bson:"url,required"`
	Tasks           []valueobject.Task `bson:"tasks,omitempty"`
}
