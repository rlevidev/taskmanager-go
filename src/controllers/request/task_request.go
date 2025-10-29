package request

type CreateTaskRequest struct {
	Title       string `json:"task_title" binding:"required,min=3,max=100"`
	Description string `json:"task_description" binding:"required,min=3,max=100"`
}
