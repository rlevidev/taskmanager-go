package services

import (
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/models"
	"gorm.io/gorm"
)

func AuthenticateUser(email, password string, db *gorm.DB) (*models.UserDomain, *resterr.RestErr) {
	var user models.UserDomain

	// Buscar usuário por email
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, resterr.NewUnauthorizedError("Credenciais inválidas")
		}
		return nil, resterr.NewInternalServerError("Erro ao buscar usuário")
	}

	// Verificar senha
	if !user.CheckPassword(password) {
		return nil, resterr.NewUnauthorizedError("Credenciais inválidas")
	}

	return &user, nil
}
