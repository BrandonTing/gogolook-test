package schema

type Task struct {
	// ID     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
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
	Tasks []Task `json:"tasks"`
}

type SetTaskResponse struct {
	Tasks []TaskWithID `json:"tasks"`
}

type UpdateTasksInput struct {
	ID     string `param:"id"`
	Name   string `json:"name,omitempty"`
	Status int    `json:"status,omitempty"`
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
