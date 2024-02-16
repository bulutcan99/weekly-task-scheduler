package entity

import (
	"fmt"
	"github.com/bulutcan99/weekly-task-scheduler/internal/infrastructure/env"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"sort"
	"strconv"
)

var (
	developerNumber = &env.Env.DevNumber
)

const (
	TotalDuration = 45
)

type Developer struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	Speed          int                `bson:"speed"`
	RemainingHours int                `bson:"reamaning_hours"`
}

func NewDevelopers() []Developer {
	devs := make([]Developer, 0)
	for i := 0; i < *developerNumber; i++ {
		devs = append(devs, Developer{
			ID:             primitive.NewObjectID(),
			Name:           "Developer-" + strconv.Itoa(i+1),
			Speed:          i + 1,
			RemainingHours: TotalDuration,
		})
	}
	sort.Slice(devs, func(i, j int) bool {
		return devs[i].Speed > devs[j].Speed
	})
	return devs
}
func (w *Developer) GetIntensityPerHour() int {
	return w.Speed
}

func (w *Developer) GetWeekleyTotalTasks() int {
	return TotalDuration * w.Speed
}

func (w *Developer) GetWeekleyTotalHours() int {
	return TotalDuration
}

func (w *Developer) Work(taskName string) {
	assign := fmt.Sprintf("Task %s assigned to %s", taskName, w.Name)
	slog.Debug(assign)
}
