package tasks

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Task struct {
	// ID     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type TaskStore map[string]Task

func CreateTaskStore() TaskStore {
	return map[string]Task{}
}

type GetTaskResponse struct {
	Tasks []Task `json:"tasks"`
}

func (t TaskStore) GetTasks(c echo.Context) error {
	tasks := []Task{}

	for _, val := range t {
		tasks = append(tasks, val)
	}

	return c.JSON(http.StatusOK, GetTaskResponse{
		Tasks: tasks,
	})
}

type SetTasksInput struct {
	Tasks []Task `json:"tasks"`
}

type SetTaskResponse struct {
	IsSuccess bool     `json:"isSuccess"`
	IDList    []string `json:"idList"`
}

func (t TaskStore) SetTasks(c echo.Context) error {
	content := SetTasksInput{}
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}
	if content.Tasks == nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}
	fmt.Printf("%v\n", content)
	ids := []string{}
	for _, task := range content.Tasks {
		// create unique id
		newUUID := uuid.New().String()
		t[newUUID] = task
		ids = append(ids, newUUID)
	}
	return c.JSON(http.StatusOK, SetTaskResponse{
		IsSuccess: true,
		IDList:    ids,
	})
}

// how to validate all key exist?
type UpdateTasksInput struct {
	ID     string `param:"id"`
	Name   string `json:"name,omitempty"`
	Status int    `json:"status,omitempty"`
}

type UpdateTaskResponse struct {
	IsSuccess bool   `json:"isSuccess"`
	ID        string `json:"id"`
}

func (t TaskStore) UpdateTasks(c echo.Context) error {
	content := UpdateTasksInput{}

	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}
	_, ok := t[content.ID]
	// If the key exists
	if !ok {
		return c.JSON(http.StatusBadRequest, "Task not found.")
	}

	t[content.ID] = Task{
		Name:   content.Name,
		Status: content.Status,
	}

	return c.JSON(http.StatusOK, UpdateTaskResponse{
		IsSuccess: true,
		ID:        content.ID,
	})
}

type RemoveTaskInput struct {
	ID string `param:"id"`
}

type RemoveTasksResponse struct {
	Name string `json:"name"`
}

func (t TaskStore) RemoveTasks(c echo.Context) error {
	content := UpdateTasksInput{}

	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}
	task, ok := t[content.ID]
	// If the key exists
	if !ok {
		return c.JSON(http.StatusBadRequest, "Task not found.")
	}
	delete(t, content.ID)
	return c.JSON(http.StatusOK, RemoveTasksResponse{
		Name: task.Name,
	})
}
