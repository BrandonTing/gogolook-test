package tests

import (
	"bytes"
	"encoding/json"
	"gogolook-test/internal/schema"
	"gogolook-test/internal/storage"
	"gogolook-test/internal/tasks"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func TestGetHandler(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	if err := tasks.GetTasksHandler(c); err != nil {
		t.Errorf("GetTasksHandler() error = %v", err)
		return
	}

	if resp.Code != http.StatusOK {
		t.Errorf("GetTasksHandler() wrong status code = %v", resp.Code)
		return
	}

	expected := schema.GetTaskResponse{
		Tasks: []schema.TaskWithID{},
	}

	var actual schema.GetTaskResponse
	// Decode the response body into the actual map
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		t.Errorf("GetTasksHandler() error decoding response body: %v", err)
		return
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("GetTasksHandler() wrong response body. expected = %v, actual = %v", expected, actual)
		return
	}
}

func TestCreateHandler(t *testing.T) {

	e := echo.New()
	e.Validator = &schema.Validator{Validator: validator.New()}

	expected := []schema.Task{
		{
			Name:   "test",
			Status: schema.GetIntPointer(0),
		},
	}
	paramsByte, _ := json.Marshal(schema.SetTasksInput{
		Tasks: expected,
	})
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(paramsByte))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	if err := tasks.SetTasksHandler(c); err != nil {
		t.Errorf("SetTasksHandler() error = %v", err)
		return
	}

	if resp.Code != http.StatusOK {
		t.Errorf("SetTasksHandler() wrong status code = %v", resp.Code)
		return
	}

	var actual schema.SetTaskResponse
	// Decode the response body into the actual map
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		t.Errorf("SetTasksHandler() error decoding response body: %v", err)
		return
	}
	if len(actual.Tasks) != len(expected) || actual.Tasks[0].Name != expected[0].Name {
		t.Errorf("Failed to create task")
		return
	}

}

func TestUpdateHandler(t *testing.T) {

	tasks.TaskStore.Data["test-id"] = storage.ItemWithID[schema.Task]{
		ID: "test-id",
		Item: schema.Task{
			Name:   "test",
			Status: schema.GetIntPointer(0),
		},
	}
	e := echo.New()
	e.Validator = &schema.Validator{Validator: validator.New()}

	paramsByte, _ := json.Marshal(schema.Task{
		Name:   "new name",
		Status: schema.GetIntPointer(0),
	})

	req := httptest.NewRequest(http.MethodPut, "/tasks/test-id", bytes.NewBuffer(paramsByte))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	c.SetParamNames("id")
	c.SetParamValues("test-id")
	if err := tasks.UpdateTasksHandler(c); err != nil {
		t.Errorf("UpdateTasksHandler() error = %v", err)
		return
	}

	if resp.Code != http.StatusOK {
		t.Errorf("UpdateTasksHandler() wrong status code = %v", resp.Code)
		return
	}
	expected := schema.UpdateTaskResponse{
		Task: schema.TaskWithID{
			ID:     "test-id",
			Name:   "new name",
			Status: 0,
		},
	}
	var actual schema.UpdateTaskResponse
	// Decode the response body into the actual map
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		t.Errorf("UpdateTasksHandler() error decoding response body: %v", err)
		return
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Failed to update task")
		return
	}

}

func TestRemoveHandler(t *testing.T) {

	tasks.TaskStore.Data["test-id"] = storage.ItemWithID[schema.Task]{
		ID: "test-id",
		Item: schema.Task{
			Name:   "test",
			Status: schema.GetIntPointer(0),
		},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/tasks/test-id", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	c.SetParamNames("id")
	c.SetParamValues("test-id")

	if err := tasks.RemoveTasksHandler(c); err != nil {
		t.Errorf("RemoveTasksHandler() error = %v", err)
		return
	}

	if resp.Code != http.StatusOK {
		t.Errorf("RemoveTasksHandler() wrong status code = %v", resp.Code)
		return
	}
	expected := schema.RemoveTasksResponse{
		Name: "test",
	}
	var actual schema.RemoveTasksResponse
	// Decode the response body into the actual map
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		t.Errorf("RemoveTasksHandler() error decoding response body: %v", err)
		return
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Failed to remove task")
		return
	}

}
