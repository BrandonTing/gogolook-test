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

func init() {
	SetupStore()
}

func (s *Store) GetByID(id string) (*schema.TaskWithID, error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	task, ok := s.Data[id]
	// If the key exists
	if !ok {
		return nil, fmt.Errorf("task not found")
	}

	return &schema.TaskWithID{
		ID:     id,
		Name:   task.Name,
		Status: *task.Status,
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
			Status: *val.Status,
		})
	}

	return tasks, nil
}

func (s *Store) Create(tasks []schema.Task) ([]schema.TaskWithID, error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	taskWithIDs := []schema.TaskWithID{}
	for _, task := range tasks {
		// create unique id
		newUUID := uuid.New().String()
		s.Data[newUUID] = task
		taskWithIDs = append(taskWithIDs, schema.TaskWithID{
			Name:   task.Name,
			Status: *task.Status,
			ID:     newUUID,
		})
	}

	return taskWithIDs, nil
}

func (s *Store) Update(param schema.UpdateTasksInput) (*schema.TaskWithID, error) {
	if s == nil || s.Data == nil {
		return nil, fmt.Errorf("store is not setup correctly")
	}
	_, ok := s.Data[param.ID]
	// If the key exists
	if !ok {
		return nil, fmt.Errorf("task not found")
	}

	newTask := schema.Task{
		Name:   param.Name,
		Status: param.Status,
	}

	s.Data[param.ID] = newTask

	return &schema.TaskWithID{
		ID:     param.ID,
		Name:   newTask.Name,
		Status: *newTask.Status,
	}, nil
}

func (s *Store) Remove(id string) (string, error) {
	if s == nil || s.Data == nil {
		return "", fmt.Errorf("store is not setup correctly")
	}
	value, ok := s.Data[id]
	// If the key exists
	if !ok {
		return "", fmt.Errorf("task not found")
	}
	delete(s.Data, id)
	return value.Name, nil
}
