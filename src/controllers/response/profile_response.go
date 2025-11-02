package response

import (
	"time"
)

type TaskResponse struct {
	TaskID    string `json:"task_id"`
	TaskTitle string `json:"task_title"`
	TaskStatus string `json:"task_status"`
}

type ProfileResponse struct {
	UserID      string          `json:"user_id"`
	UserName    string          `json:"user_name"`
	UserCreateAt time.Time      `json:"user_create_at"`
	Tasks       []TaskResponse `json:"tasks"`
}
