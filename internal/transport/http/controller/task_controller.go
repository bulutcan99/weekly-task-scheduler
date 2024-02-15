package controller

import (
	"fmt"
	"github.com/bulutcan99/weekly-task-scheduler/internal/application/interfaces"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/aggregate"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/valueobject"
	"github.com/gofiber/fiber/v3"
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

type TaskChQueue struct {
	Id   int64
	Task valueobject.Task
}

func (tc *TaskController) AssignTask(ctx fiber.Ctx) error {
	developers := entity.NewDevelopers()
	tasks, err := tc.TaskService.GetTasks(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error while getting tasks",
		})
	}
	devTasks := make(map[string]*aggregate.DeveloperTask)

	poolSize := len(developers)
	taskQueue := make(chan TaskChQueue, len(tasks))
	resultQueue := make(chan *aggregate.DeveloperTask, len(tasks))
	var wg sync.WaitGroup
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go worker(&wg, i+1, &developers, taskQueue, resultQueue)
	}

	for i := 0; i < len(tasks); i++ {
		taskQueue <- TaskChQueue{Id: int64(i), Task: tasks[i]}
	}

	close(taskQueue)
	wg.Wait()
	close(resultQueue)

	for result := range resultQueue {
		devTasks[result.Name] = result
	}

	return ctx.JSON(devTasks)
}

func worker(wg *sync.WaitGroup, id int, developers *[]entity.Developer, taskQueue <-chan TaskChQueue, resultQueue chan<- *aggregate.DeveloperTask) {
	defer wg.Done()

	for task := range taskQueue {
		devTask := &aggregate.DeveloperTask{Name: fmt.Sprintf("Developer-%d", id), Weeks: 0}
		scheduleTask(developers, task.Task, devTask)
		resultQueue <- devTask
	}
}

func scheduleTask(developers *[]entity.Developer, task valueobject.Task, devTask *aggregate.DeveloperTask) {
	for _, dev := range *developers {
		remainingWork := task.Difficulty - dev.WorkedHours
		workThisWeek := min(remainingWork, dev.WeeklyTotalHours)
		dev.WorkedHours += workThisWeek
		dev.WeeklyTotalHours -= workThisWeek
		task.Difficulty -= workThisWeek

		devTask.Tasks = append(devTask.Tasks, valueobject.Task{TaskName: task.TaskName, Duration: task.Duration, Difficulty: workThisWeek})
		devTask.Total += workThisWeek

		if task.Difficulty == 0 {
			break
		}
	}

	if devTask.Total > 0 {
		devTask.Weeks++
		devTask.Remain = task.Difficulty
	}
}
