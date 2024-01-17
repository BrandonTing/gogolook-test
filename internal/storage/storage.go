package storage

import (
	"fmt"
	"gogolook-test/internal/schema"

	"github.com/google/uuid"
)

type Store struct {
	Data map[string]schema.Task
}

var TaskStore *Store

func SetupStore() {
	TaskStore = &Store{
		Data: map[string]schema.Task{},
	}
}

func (s *Store) GetByID(id string) (*schema.TaskWithID, error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	task, ok := s.Data[id]
	// If the key exists
	if !ok {
		return nil, fmt.Errorf("Task not found.")
	}

	return &schema.TaskWithID{
		ID:     id,
		Name:   task.Name,
		Status: task.Status,
	}, nil
}

func (s *Store) GetAll() ([]schema.TaskWithID, error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	tasks := []schema.TaskWithID{}

	for key, val := range s.Data {
		tasks = append(tasks, schema.TaskWithID{
			ID:     key,
			Name:   val.Name,
			Status: val.Status,
		})
	}

	return tasks, nil
}

func (s *Store) Create(param schema.SetTasksInput) ([]schema.TaskWithID, error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	tasks := []schema.TaskWithID{}
	for _, task := range param.Tasks {
		// create unique id
		newUUID := uuid.New().String()
		fmt.Printf("uuid %v", newUUID)
		s.Data[newUUID] = task
		fmt.Printf("uuid %v", newUUID)
		tasks = append(tasks, schema.TaskWithID{
			Name:   task.Name,
			Status: task.Status,
			ID:     newUUID,
		})
	}

	return tasks, nil
}

func (s *Store) Update(param schema.UpdateTasksInput) (*schema.TaskWithID, error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	_, ok := s.Data[param.ID]
	// If the key exists
	if !ok {
		return nil, fmt.Errorf("Task not found.")
	}

	newTask := schema.Task{
		Name:   param.Name,
		Status: param.Status,
	}

	s.Data[param.ID] = newTask

	return &schema.TaskWithID{
		ID:     param.ID,
		Name:   newTask.Name,
		Status: newTask.Status,
	}, nil
}

func (s *Store) Remove(param schema.RemoveTaskInput) (string, error) {
	if s == nil || s.Data == nil {
		return "", fmt.Errorf("store is not setup correctly")
	}
	value, ok := s.Data[param.ID]
	// If the key exists
	if !ok {
		return "", fmt.Errorf("Task not found.")
	}
	delete(s.Data, param.ID)
	return value.Name, nil
}
