package valueobject

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Task struct {
	ID            primitive.ObjectID `bson:"_id"`
	ProviderID    primitive.ObjectID `bson:"provider_id"`
	DeveloperName string             `bson:"developer_name"`
	TaskName      string             `bson:"task_name"`
	Difficulty    int                `bson:"difficulty"`
	Duration      int                `bson:"duration"`
	Intensity     int                `bson:"intensity"`
	CreatedAt     time.Time          `bson:"created_at"`
}

func (t *Task) GetIntensity() int {
	return t.Difficulty * t.Duration
}
