package tasks

import (
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTasksHandler(c echo.Context) error {
	tasks, err := storage.TaskStore.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve tasks")
	}
	return c.JSON(http.StatusOK, schema.GetTaskResponse{
		Tasks: tasks,
	})
}

func SetTasksHandler(c echo.Context) error {

	content := schema.SetTasksInput{}
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}
	if content.Tasks == nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}

	tasks, err := storage.TaskStore.Create(content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create tasks")
	}

	return c.JSON(http.StatusOK, schema.SetTaskResponse{
		IsSuccess: true,
		Tasks:     tasks,
	})
}

// TODO how to validate all key exist?

func UpdateTasksHandler(c echo.Context) error {
	content := schema.UpdateTasksInput{}

	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}

	task, err := storage.TaskStore.Update(content)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}

	return c.JSON(http.StatusOK, schema.UpdateTaskResponse{
		IsSuccess: true,
		Task:      *task,
	})
}

func RemoveTasksHandler(c echo.Context) error {
	content := schema.RemoveTaskInput{}

	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}
	name, err := storage.TaskStore.Remove(content)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}
	return c.JSON(http.StatusOK, schema.RemoveTasksResponse{
		Name: *name,
	})
}
