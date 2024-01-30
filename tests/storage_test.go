package tests

import (
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"reflect"
	"testing"
)

var TaskStore *storage.Store[schema.Task]

func init() {
	TaskStore = storage.SetupStore[schema.Task]()
}

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
	tasks, err := TaskStore.GetAll()
	if err != nil {
		t.Errorf("Failed to execute GetAll method: %v", err)
		return
	}
	expected := []storage.ItemWithID[schema.Task]{}
	if !reflect.DeepEqual(expected, tasks) {
		t.Errorf("initial store should contain 0 item but get %v", tasks)
		return
	}
}

func TestCreate(t *testing.T) {
	newTasks := CreateTestTasks()
	newTaskMap := make(map[string]int)
	for _, task := range newTasks {
		newTaskMap[task.Name] = *task.Status
	}
	tasks, err := TaskStore.Create(newTasks)
	if err != nil {
		t.Errorf("Failed to execute Create method: %v", err)
		return
	}
	// test if correctly created ID
	for _, task := range tasks {
		if _, ok := newTaskMap[task.Item.Name]; !ok {
			t.Errorf("Failed to create target task: %v", err)
			return
		}
		if *task.Item.Status != newTaskMap[task.Item.Name] {
			t.Errorf("Failed to init task with correct status: %v", err)
			return
		}
		if task.ID == "" {
			t.Errorf("Failed to create unique ID: %v", err)
			return
		}
	}
	// test get all after create
	fetchedTasks, _ := TaskStore.GetAll()
	if !reflect.DeepEqual(fetchedTasks, tasks) {
		t.Errorf("Failed to get tasks just created: %v", err)
		return
	}
}

func TestGetByID(t *testing.T) {
	newTasks := CreateTestTasks()
	tasks, _ := TaskStore.Create(newTasks)
	taskToCheck := tasks[0]
	actualTask, err := TaskStore.GetByID(taskToCheck.ID)
	if err != nil {
		t.Errorf("Failed to execute GetByID: %v", err)
		return
	}
	if actualTask.Item.Name != taskToCheck.Item.Name || actualTask.Item.Status != taskToCheck.Item.Status {
		t.Errorf("Failed to retrieve target task: %v", err)
		return
	}
}

func TestUpdate(t *testing.T) {
	newTasks := CreateTestTasks()
	newTaskMap := map[string]int{}
	for _, task := range newTasks {
		newTaskMap[task.Name] = *task.Status
	}
	tasks, _ := TaskStore.Create(newTasks)
	taskToUpdate := tasks[0]
	_, err := TaskStore.Update(storage.ItemWithID[schema.Task]{
		ID: taskToUpdate.ID,
		Item: schema.Task{
			Name:   "new name",
			Status: schema.GetIntPointer(1),
		},
	})
	if err != nil {
		t.Errorf("Failed to update task: %v", err)
		return
	}
	newTask, err := TaskStore.GetByID(taskToUpdate.ID)
	if err != nil {
		t.Errorf("Failed to get target task: %v", err)
		return
	}
	if newTask.Item.Name != "new name" || *newTask.Item.Status != 1 {
		t.Errorf("Failed to update task correctly: %v", err)
		return
	}
}
func TestRemove(t *testing.T) {
	newTasks := CreateTestTasks()
	tasks, _ := TaskStore.Create(newTasks)
	taskToRemove := tasks[0]
	removed, err := TaskStore.Remove(taskToRemove.ID)
	if err != nil {
		t.Errorf("Failed to remove task: %v", err)
		return
	}
	if removed.Item.Name != taskToRemove.Item.Name {
		t.Errorf("Failed to remove correct task: should be %v but deleted %v", taskToRemove.Item.Name, removed.Item.Name)
		return
	}
	if _, err := TaskStore.GetByID(taskToRemove.ID); err == nil {
		t.Errorf("Target task still exist: %v", err)
		return
	}
}
