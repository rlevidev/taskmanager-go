package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/taskmanager-go/src/config/validation"
	"github.com/rlevidev/taskmanager-go/src/controllers/request"
	"github.com/rlevidev/taskmanager-go/src/controllers/response"
	"github.com/rlevidev/taskmanager-go/src/services"
	"gorm.io/gorm"
)

func LoginUser(ctx *gin.Context, db *gorm.DB) {
	var userRequest request.UserLoginRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Status, restErr)
		return
	}

	// Autenticar usuário
	user, err := services.AuthenticateUser(userRequest.Email, userRequest.Password, db)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	// Gerar token JWT
	token, err := services.GenerateToken(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	// Retornar resposta
	userResponse := response.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	ctx.JSON(200, gin.H{
		"message": "Login realizado com sucesso",
		"data": gin.H{
			"user":  userResponse,
			"token": token,
		},
	})
}
