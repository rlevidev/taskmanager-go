package services

import (
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/models"
	"gorm.io/gorm"
)

func CreateTask(td models.TaskDomain, db *gorm.DB) (*models.TaskDomain, *resterr.RestErr) {
	if err := db.Create(&td).Error; err != nil {
		return nil, resterr.NewInternalServerError("Erro ao salvar tarefa no banco de dados")
	}

	return &td, nil
}
