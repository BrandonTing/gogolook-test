package web

import (
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"net/http"
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

func TasksHomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := TaskStore.GetAll()
	if err != nil {
		http.Error(w, "Failed to render list", http.StatusInternalServerError)
		return
	}
	tasks := fromDataListToTasks(data)
	component := TasksHome(tasks)
	component.Render(r.Context(), w)
}

func NewTaskHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	_, err = TaskStore.Create([]schema.Task{
		{
			Name:   name,
			Status: schema.GetIntPointer(0),
		},
	})
	if err != nil {
		http.Error(w, "Failed to create new task", http.StatusInternalServerError)
		return
	}
	data, err := TaskStore.GetAll()
	if err != nil {
		http.Error(w, "Failed to get current list", http.StatusInternalServerError)
		return
	}
	tasks := fromDataListToTasks(data)
	component := TasksTable(tasks)
	component.Render(r.Context(), w)
}

func UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")
	data, err := TaskStore.GetByID(id)
	if err != nil {
		http.Error(w, "Failed to get task", http.StatusNotFound)
		return
	}
	task := fromDataToTask(*data)
	newStatus := 0
	if task.Status == 0 {
		newStatus = 1
	}
	_, err = TaskStore.Update(storage.ItemWithID[schema.Task]{
		ID: id,
		Item: schema.Task{
			Name:   task.Name,
			Status: &newStatus,
		},
	})
	if err != nil {
		http.Error(w, "Failed to update task status", http.StatusInternalServerError)
		return
	}
	component := TaskStatus(id, newStatus)
	component.Render(r.Context(), w)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	_, err := TaskStore.Remove(id)
	if err != nil {
		http.Error(w, "Failed to remove task", http.StatusInternalServerError)
		return
	}
}
