package models

import (
	"time"

	"github.com/google/uuid"
)

type TaskDomain struct {
	ID          string    `json:"task_id" gorm:"primaryKey"`
	Title       string    `json:"task_title" gorm:"not null"`
	Description string    `json:"task_description" gorm:"not null"`
	Status      string    `json:"task_status" gorm:"default:'pending'"`
	UserID      string    `json:"user_id" gorm:"not null"`
	CreatedAt   time.Time `json:"task_created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewTaskDomain(
	Title string,
	Description string,
	UserID string,
) *TaskDomain {
	return &TaskDomain{
		ID:          uuid.New().String(),
		Title:       Title,
		Description: Description,
		Status:      "pending",
		UserID:      UserID,
	}
}
