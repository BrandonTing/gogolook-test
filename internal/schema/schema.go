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

type GetTaskResponse struct {
	Tasks []TaskWithID `json:"tasks"`
}

type SetTasksInput struct {
	Tasks []Task `json:"tasks"`
}

type SetTaskResponse struct {
	IsSuccess bool         `json:"isSuccess"`
	Tasks     []TaskWithID `json:"tasks"`
}

type UpdateTasksInput struct {
	ID     string `param:"id"`
	Name   string `json:"name,omitempty"`
	Status int    `json:"status,omitempty"`
}

type UpdateTaskResponse struct {
	IsSuccess bool       `json:"isSuccess"`
	Task      TaskWithID `json:"task"`
}

type RemoveTaskInput struct {
	ID string `param:"id"`
}

type RemoveTasksResponse struct {
	Name string `json:"name"`
}
