package services

import (
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/models"
	"gorm.io/gorm"
)

func CreateUser(ud models.UserDomain, db *gorm.DB) (*models.UserDomain, *resterr.RestErr) {
	if err := ud.EncryptPassword(); err != nil {
		return nil, resterr.NewInternalServerError("Erro ao encriptar senha")
	}

	if err := db.Create(&ud).Error; err != nil {
		return nil, resterr.NewInternalServerError("Erro ao salvar usuário no banco de dados")
	}

	return &ud, nil
}
