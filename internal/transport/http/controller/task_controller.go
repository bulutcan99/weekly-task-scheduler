package controller

import (
	"fmt"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/gofiber/fiber/v3"
	"math"
	"sync"
)

type TaskController struct {
	TaskService interfaces.ITaskService
}

func NewTaskController(taskService interfaces.ITaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

func (tc *TaskController) GetTasks(ctx fiber.Ctx) error {
	tasks, err := tc.TaskService.GetTasks(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error while getting tasks",
		})
	}
	return ctx.JSON(tasks)

}

// Bunlar map ile donulcek (aggregate data ile) ve swagger halledilcek, bide description halledilcek. Bide task upsert kisminda insert ederken eski data gelince hata veriyo.
func (tc *TaskController) AssignTask(ctx fiber.Ctx) error {
	developers := entity.NewDevelopers()
	tasks, err := tc.TaskService.GetTasks(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error while getting tasks",
		})
	}
	weeks := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for {
		weeks++
		for _, developer := range developers {
			wg.Add(1)
			go func(dev *entity.Developer) {
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				hoursLeft := dev.RemainingHours
				for i := len(tasks) - 1; i >= 0; i-- {
					task := tasks[i]
					taskEfficiency := task.Difficulty * task.Duration
					taskHours := math.Ceil(float64(taskEfficiency) / float64(dev.Speed))
					if taskHours <= float64(hoursLeft) {
						fmt.Printf("Week %d\n", weeks)
						dev.Work(task.Name)
						hoursLeft -= int(taskHours)
						tasks = append(tasks[:i], tasks[i+1:]...)
					}
				}
			}(&developer)
		}
		wg.Wait()
		if len(tasks) == 0 {
			break
		}
	}

	fmt.Printf("Total weeks to complete all tasks: %d\n", weeks)
	return ctx.JSON(developers)
}
