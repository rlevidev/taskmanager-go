package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/taskmanager-go/src/config/validation"
	"github.com/rlevidev/taskmanager-go/src/controllers/request"
	"gorm.io/gorm"
)

func CreateUser(ctx *gin.Context, db *gorm.DB) {
	var userRequest request.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Status, restErr)
		return
	}

	ctx.JSON(201, gin.H{
		"message": "User created successfully",
		"data":    userRequest,
	})
}
