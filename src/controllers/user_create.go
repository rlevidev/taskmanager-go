package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/config/validation"
	"github.com/rlevidev/taskmanager-go/src/controllers/request"
	"github.com/rlevidev/taskmanager-go/src/controllers/response"
	"github.com/rlevidev/taskmanager-go/src/models"
	"github.com/rlevidev/taskmanager-go/src/services"
	"gorm.io/gorm"
)

func CreateUser(ctx *gin.Context, db *gorm.DB) {
	var userRequest request.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Status, restErr)
		return
	}

	// Verificar se email já existe
	var existingUser models.UserDomain
	if err := db.Where("email = ?", userRequest.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(400, resterr.NewBadRequestError("Email já está em uso"))
		return
	} else if err != gorm.ErrRecordNotFound {
		ctx.JSON(500, resterr.NewInternalServerError("Erro ao verificar email"))
		return
	}

	// Criar novo usuário
	userDomain := models.NewUserDomain(
		userRequest.Email,
		userRequest.Name,
		userRequest.Password,
	)

	// Salvar no banco
	createdUser, err := services.CreateUser(*userDomain, db)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	// Retornar resposta
	userResponse := response.UserResponse{
		ID:    createdUser.ID,
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}

	ctx.JSON(201, gin.H{
		"message": "Usuário criado com sucesso",
		"data":    userResponse,
	})
}
