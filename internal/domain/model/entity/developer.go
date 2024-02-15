package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
	"strconv"
)

const (
	TotalDuration = 5
)

type Developer struct {
	ID               primitive.ObjectID `bson:"_id"`
	Name             string             `bson:"name"`
	Speed            int                `bson:"speed"`
	WorkedHours      int                `bson:"worked_hours"`
	WeeklyTotalHours int                `bson:"total_hours"`
}

func NewDevelopers() []Developer {
	devs := make([]Developer, 0)
	for i := 0; i < 5; i++ {
		devs = append(devs, Developer{
			ID:               primitive.NewObjectID(),
			Name:             "Developer-" + strconv.Itoa(i+1),
			Speed:            i + 1,
			WorkedHours:      0,
			WeeklyTotalHours: TotalDuration,
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

func (w *Developer) Work(duration int) {
	w.WorkedHours += duration
}
