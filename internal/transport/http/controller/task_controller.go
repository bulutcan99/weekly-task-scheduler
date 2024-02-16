package controller

import (
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/dto"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/aggregate"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"github.com/gofiber/fiber/v2"
	"math"
	"sync"
)

type TaskController struct {
	TaskService interfaces.ITaskService
}

func NewTaskController(taskService interfaces.ITaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

// @Summary Get all the tasks
// @Description Get all the task from the database
// @ID get-tasks
// @Produce json
// @Success 200 {object} []dto.Task
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/providers [get]
func (tc *TaskController) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := tc.TaskService.GetTasks(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(dto.ErrorResponse{
			Error: true,
			Msg:   "Error while getting tasks",
		})
	}
	return ctx.JSON(tasks)

}

// Bunlar map ile donulcek (aggregate data ile) ve swagger halledilcek, bide description halledilcek. Bide task upsert kisminda insert ederken eski data gelince hata veriyo.

// @Summary Assign tasks to developers
// @Description Assign tasks to developers
// @ID assign-tasks
// @Produce json
// @Success 200 {object} map[string]aggregate.DeveloperTask
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/task [post]
func (tc *TaskController) AssignTask(ctx *fiber.Ctx) error {
	developers := entity.NewDevelopers()
	tasks, err := tc.TaskService.GetTasks(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(dto.ErrorResponse{
			Error: true,
			Msg:   "Error while getting tasks",
		})
	}
	devMap := make(map[string]*aggregate.DeveloperTask)
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
					taskEfficiency := task.Intensity
					taskHours := math.Ceil(float64(taskEfficiency) / float64(dev.Speed))

					if taskHours <= float64(hoursLeft) {
						dev.Work(task.Name)
						if devMap[dev.Name] == nil {
							devMap[dev.Name] = &aggregate.DeveloperTask{
								Developer: dev,
								Tasks:     &[]valueobject.Task{},
							}
						}

						*devMap[dev.Name].Tasks = append(*devMap[dev.Name].Tasks, task)
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

	for _, devTask := range devMap {
		devTask.Week = weeks
	}

	return ctx.Status(fiber.StatusOK).JSON(devMap)
}
