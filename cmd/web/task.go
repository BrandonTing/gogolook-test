package web

import (
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"net/http"
)

func TasksHomeHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := storage.TaskStore.GetAll()
	if err != nil {
		http.Error(w, "Failed to render list", http.StatusInternalServerError)
		return
	}

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
	_, err = storage.TaskStore.Create([]schema.Task{
		{
			Name:   name,
			Status: schema.GetIntPointer(0),
		},
	})
	if err != nil {
		http.Error(w, "Failed to create new task", http.StatusInternalServerError)
		return
	}
	tasks, err := storage.TaskStore.GetAll()
	if err != nil {
		http.Error(w, "Failed to get current list", http.StatusInternalServerError)
		return
	}
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
	task, err := storage.TaskStore.GetByID(id)
	if err != nil {
		http.Error(w, "Failed to get task", http.StatusNotFound)
		return
	}

	newStatus := 0
	if task.Status == 0 {
		newStatus = 1
	}
	_, err = storage.TaskStore.Update(schema.UpdateTasksInput{
		ID:     id,
		Name:   task.Name,
		Status: &newStatus,
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
	_, err := storage.TaskStore.Remove(id)
	if err != nil {
		http.Error(w, "Failed to remove task", http.StatusInternalServerError)
		return
	}
}
