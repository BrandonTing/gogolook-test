package tasks

import (
	"fmt"
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTasksHandler(c echo.Context) error {
	tasks, err := storage.TaskStore.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.FailResponse{Message: "Failed to retrieve tasks"})
	}
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
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "invalid input"})
	}

	tasks, err := storage.TaskStore.Create(content)
	if err != nil {
		fmt.Printf("[Create task] error: %v\n", err)
		return c.JSON(http.StatusInternalServerError, schema.FailResponse{Message: "Failed to create tasks"})
	}

	return c.JSON(http.StatusOK, schema.SetTaskResponse{
		Tasks: tasks,
	})
}

func UpdateTasksHandler(c echo.Context) error {
	content := schema.UpdateTasksInput{}

	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "invalid input"})
	}

	task, err := storage.TaskStore.Update(content)
	if err != nil {
		fmt.Printf("[Update task] error: %v\n", err)
		return c.JSON(http.StatusNotFound, schema.FailResponse{Message: "Failed to update Task"})
	}

	return c.JSON(http.StatusOK, schema.UpdateTaskResponse{
		Task: *task,
	})
}

func RemoveTasksHandler(c echo.Context) error {
	content := schema.RemoveTaskInput{}
	fmt.Printf("%v", content)
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, schema.FailResponse{Message: "invalid input"})
	}
	name, err := storage.TaskStore.Remove(content)
	if err != nil {
		fmt.Printf("[Remove task] error: %v\n", err)
		return c.JSON(http.StatusNotFound, schema.FailResponse{Message: "Failed to remove task"})
	}
	return c.JSON(http.StatusOK, schema.RemoveTasksResponse{
		Name: name,
	})
}
