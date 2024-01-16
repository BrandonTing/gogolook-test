package tasks

import (
	"github.com/google/uuid"
)

type Status int

const (
	incomplete Status = iota
	completed  Status = 1
)

type Task struct {
	// ID     string `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

type TaskStore map[string]Task

func CreateTaskStore() TaskStore {
	return map[string]Task{}
}

func (t *TaskStore) GetTasks() []Task {
	tasks := []Task{}

	for _, val := range *t {
		tasks = append(tasks, val)
	}

	return tasks
}

func (t *TaskStore) SetTasks(data []Task) error {
	for _, task := range data {
		// create unique id
		newUUID := uuid.New()
		(*t)[newUUID.String()] = task
	}
	return nil
}

func (t *TaskStore) UpdateTasks(id string, task Task) error {
	tasks := []Task{}

	for _, val := range *t {
		tasks = append(tasks, val)
	}
	return nil
}

func (t *TaskStore) RemoveTasks(id string) error {
	tasks := []Task{}

	for _, val := range *t {
		tasks = append(tasks, val)
	}
	return nil
}
