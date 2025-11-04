package services

import (
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/models"
	"gorm.io/gorm"
)

func UpdateTaskStatus(taskID string, userID string, status string, db *gorm.DB) *resterr.RestErr {
	var task models.TaskDomain

	// Verificar se a tarefa existe e pertence ao usuário
	if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return resterr.NewNotFoundError("Tarefa não encontrada")
		}
		return resterr.NewInternalServerError("Erro ao buscar tarefa")
	}

	// Atualizar o status
	if err := db.Model(&task).Update("status", status).Error; err != nil {
		return resterr.NewInternalServerError("Erro ao atualizar status da tarefa")
	}

	return nil
}

func DeleteTask(taskID string, userID string, db *gorm.DB) *resterr.RestErr {
	var task models.TaskDomain

	// Verificar se a tarefa existe e pertence ao usuário
	if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return resterr.NewNotFoundError("Tarefa não encontrada")
		}
		return resterr.NewInternalServerError("Erro ao buscar tarefa")
	}

	// Deletar a tarefa
	if err := db.Delete(&task).Error; err != nil {
		return resterr.NewInternalServerError("Erro ao deletar tarefa")
	}

	return nil
}
