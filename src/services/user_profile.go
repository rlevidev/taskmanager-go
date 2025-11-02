package services

import (
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/controllers/response"
	"github.com/rlevidev/taskmanager-go/src/models"
	"gorm.io/gorm"
)

func GetUserProfile(userID string, db *gorm.DB) (*response.ProfileResponse, *resterr.RestErr) {
	var user models.UserDomain

	// Fetch user by ID
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, resterr.NewNotFoundError("Usuário não encontrado")
		}
		return nil, resterr.NewInternalServerError("Erro ao buscar usuário")
	}

	var tasks []models.TaskDomain

	// Fetch tasks for the user
	if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, resterr.NewInternalServerError("Erro ao buscar tarefas")
	}

	// Build response
	profile := &response.ProfileResponse{
		UserID:       user.ID,
		UserName:     user.Name,
		UserCreateAt: user.CreatedAt,
		Tasks:        make([]response.TaskResponse, len(tasks)),
	}

	for i, task := range tasks {
		profile.Tasks[i] = response.TaskResponse{
			TaskID:     task.ID,
			TaskTitle:  task.Title,
			TaskStatus: task.Status,
		}
	}

	return profile, nil
}
