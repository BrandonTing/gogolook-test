package tasks

import (
	"fmt"
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

var TaskStore *storage.Store[schema.Task]

func init() {
	TaskStore = storage.SetupStore[schema.Task]()
}

func fromDataToTask(data storage.ItemWithID[schema.Task]) schema.TaskWithID {
	return schema.TaskWithID{
		ID:     data.ID,
		Name:   data.Item.Name,
		Status: *data.Item.Status,
	}
}

func fromDataListToTasks(data []storage.ItemWithID[schema.Task]) []schema.TaskWithID {
	tasks := []schema.TaskWithID{}
	for _, task := range data {
		tasks = append(tasks, fromDataToTask(task))
	}
	return tasks
}

func GetTasksHandler(c echo.Context) error {
	data, err := TaskStore.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.FailResponse{Message: "Failed to retrieve tasks"})
	}
	tasks := fromDataListToTasks(data)
	return c.JSON(http.StatusOK, schema.GetTaskResponse{
		Tasks: tasks,
	})
}

func SetTasksHandler(c echo.Context) error {

	content := schema.SetTasksInput{}
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "invalid input"})
	}
	if content.Tasks == nil {
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "Please provide at least one task"})
	}
	if err := c.Validate(content); err != nil {
		fmt.Printf("validation error: %v", err)
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "invalid input"})
	}

	data, err := TaskStore.Create(content.Tasks)
	if err != nil {
		fmt.Printf("[Create task] error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, schema.FailResponse{Message: "Failed to create tasks"})
	}
	tasks := fromDataListToTasks(data)

	return c.JSON(http.StatusOK, schema.SetTaskResponse{
		Tasks: tasks,
	})
}

func UpdateTasksHandler(c echo.Context) error {
	content := schema.UpdateTasksInput{}

	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "invalid input"})
	}
	if err := c.Validate(content); err != nil {
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "invalid input"})
	}
	_, err := TaskStore.GetByID(content.ID)
	if err != nil {
		fmt.Printf("[Update task] error: %v\n", err)
		return c.JSON(http.StatusNotFound, schema.FailResponse{Message: fmt.Sprintf("Can't find target task with ID: %v", content.ID)})
	}

	updatedTask, err := TaskStore.Update(storage.ItemWithID[schema.Task]{
		Item: schema.Task{
			Name:   content.Name,
			Status: content.Status,
		},
		ID: content.ID,
	})
	if err != nil {
		fmt.Printf("[Update task] error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, schema.FailResponse{Message: "Failed to update Task"})
	}
	return c.JSON(http.StatusOK, schema.UpdateTaskResponse{
		Task: fromDataToTask(*updatedTask),
	})
}

func RemoveTasksHandler(c echo.Context) error {
	content := schema.RemoveTaskInput{}
	fmt.Printf("%v", content)
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "invalid input"})
	}
	_, err := TaskStore.GetByID(content.ID)
	if err != nil {
		fmt.Printf("[Remove task] error: %v\n", err)
		return c.JSON(http.StatusNotFound, schema.FailResponse{Message: fmt.Sprintf("Can't find target task with ID: %v", content.ID)})
	}
	item, err := TaskStore.Remove(content.ID)
	if err != nil {
		fmt.Printf("[Remove task] error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, schema.FailResponse{Message: "Failed to remove task"})
	}
	return c.JSON(http.StatusOK, schema.RemoveTasksResponse{
		Name: item.Item.Name,
	})
}
