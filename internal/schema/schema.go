package schema

import (
	"github.com/go-playground/validator"
)

type Validator struct {
	Validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

type Task struct {
	// ID     string `json:"id"`
	Name   string `json:"name" validate:"required"`
	Status *int   `json:"status" validate:"required"`
}

type TaskWithID struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type FailResponse struct {
	Message string `json:"message"`
}

type GetTaskResponse struct {
	Tasks []TaskWithID `json:"tasks"`
}

type SetTasksInput struct {
	Tasks []Task `json:"tasks" validate:"required,dive"`
}

type SetTaskResponse struct {
	Tasks []TaskWithID `json:"tasks"`
}

type UpdateTasksInput struct {
	ID     string `param:"id"`
	Name   string `json:"name" validate:"required"`
	Status *int   `json:"status" validate:"required,oneof=0 1"`
}

type UpdateTaskResponse struct {
	Task TaskWithID `json:"task"`
}

type RemoveTaskInput struct {
	ID string `param:"id"`
}

type RemoveTasksResponse struct {
	Name string `json:"name"`
}
