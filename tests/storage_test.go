package tests

import (
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"reflect"
	"testing"
)

func CreateTestTasks() []schema.Task {
	return []schema.Task{
		{
			Name:   "test 1",
			Status: schema.GetIntPointer(0),
		},
		{
			Name:   "test 2",
			Status: schema.GetIntPointer(1),
		},
	}
}

func TestGetAll(t *testing.T) {
	storage.SetupStore()
	tasks, err := storage.TaskStore.GetAll()
	if err != nil {
		t.Errorf("Failed to execute GetAll method: %v", err)
		return
	}
	expected := []schema.TaskWithID{}
	if !reflect.DeepEqual(expected, tasks) {
		t.Errorf("initial store should contain 0 item but get %v", tasks)
		return
	}
}

func TestCreate(t *testing.T) {
	storage.SetupStore()
	newTasks := CreateTestTasks()
	newTaskMap := map[string]int{}
	for _, task := range newTasks {
		newTaskMap[task.Name] = *task.Status
	}
	tasks, err := storage.TaskStore.Create(newTasks)
	if err != nil {
		t.Errorf("Failed to execute Create method: %v", err)
		return
	}
	// test if correctly created ID
	for _, task := range tasks {
		if _, ok := newTaskMap[task.Name]; !ok {
			t.Errorf("Failed to create target task: %v", err)
			return
		}
		if task.Status != newTaskMap[task.Name] {
			t.Errorf("Failed to init task with correct status: %v", err)
			return
		}
		if task.ID == "" {
			t.Errorf("Failed to create unique ID: %v", err)
			return
		}
	}
	// test get all after create
	fetchedTasks, _ := storage.TaskStore.GetAll()
	if !reflect.DeepEqual(fetchedTasks, tasks) {
		t.Errorf("Failed to get tasks just created: %v", err)
		return
	}
}

func TestGetByID(t *testing.T) {
	storage.SetupStore()
	newTasks := CreateTestTasks()
	tasks, _ := storage.TaskStore.Create(newTasks)
	taskToCheck := tasks[0]
	actualTask, err := storage.TaskStore.GetByID(taskToCheck.ID)
	if err != nil {
		t.Errorf("Failed to execute GetByID: %v", err)
		return
	}
	if actualTask.Name != taskToCheck.Name || actualTask.Status != taskToCheck.Status {
		t.Errorf("Failed to retrieve target task: %v", err)
		return
	}
}

func TestUpdate(t *testing.T) {
	storage.SetupStore()
	newTasks := CreateTestTasks()
	newTaskMap := map[string]int{}
	for _, task := range newTasks {
		newTaskMap[task.Name] = *task.Status
	}
	tasks, _ := storage.TaskStore.Create(newTasks)
	taskToUpdate := tasks[0]
	_, err := storage.TaskStore.Update(schema.UpdateTasksInput{
		ID:     taskToUpdate.ID,
		Name:   "new name",
		Status: schema.GetIntPointer(1),
	})
	if err != nil {
		t.Errorf("Failed to update task: %v", err)
		return
	}
	newTask, err := storage.TaskStore.GetByID(taskToUpdate.ID)
	if err != nil {
		t.Errorf("Failed to get target task: %v", err)
		return
	}
	if newTask.Name != "new name" || newTask.Status != 1 {
		t.Errorf("Failed to update task correctly: %v", err)
		return
	}
}
func TestRemove(t *testing.T) {
	storage.SetupStore()
	newTasks := CreateTestTasks()
	tasks, _ := storage.TaskStore.Create(newTasks)
	taskToRemove := tasks[0]
	name, err := storage.TaskStore.Remove(taskToRemove.ID)
	if err != nil {
		t.Errorf("Failed to remove task: %v", err)
		return
	}
	if name != taskToRemove.Name {
		t.Errorf("Failed to remove correct task: should be %v but deleted %v", taskToRemove.Name, name)
		return
	}
	if _, err := storage.TaskStore.GetByID(taskToRemove.ID); err == nil {
		t.Errorf("Target task still exist: %v", err)
		return
	}
}
