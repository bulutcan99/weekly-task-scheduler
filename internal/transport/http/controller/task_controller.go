package controller

import (
	"fmt"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/gofiber/fiber/v3"
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
func (tc *TaskController) AssignTask(ctx fiber.Ctx) error {
	developers := entity.NewDevelopers()
	tasks, err := tc.TaskService.GetTasks(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error while getting tasks",
		})
	}

	assignments := make(map[string]map[string][]string)
	week := 1

	for len(tasks) > 0 {
		assignments[fmt.Sprintf("Week%d", week)] = make(map[string][]string)
		// break will gonna take for the tasks loop
		for key := len(tasks) - 1; key >= 0; key-- {
			task := tasks[key]
			for _, developer := range developers {
				taskDurationForDeveloper := task.GetIntensity() / developer.Speed
				if developer.AvailableHours >= taskDurationForDeveloper {
					developer.AvailableHours -= taskDurationForDeveloper
					assignments[fmt.Sprintf("Week%d", week)][developer.Name] = append(assignments[fmt.Sprintf("Week%d", week)][developer.Name], task.TaskName)
					tasks = append(tasks[:key], tasks[key+1:]...)
					break
				}
			}
		}

		week++
	}

	data := map[string]interface{}{
		"assignments": assignments,
		"total_weeks": week,
	}
	return ctx.JSON(data)
}
