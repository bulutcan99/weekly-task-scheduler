package valueobject

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID         primitive.ObjectID `bson:"_id"`
	ProviderID primitive.ObjectID `bson:"provider_id"`
	Name       string             `bson:"name"`
	Difficulty int                `bson:"difficulty"`
	Duration   int                `bson:"duration"`
	Intensity  int                `bson:"intensity"`
}

func (t *Task) GetIntensity() int {
	return t.Difficulty * t.Duration
}
