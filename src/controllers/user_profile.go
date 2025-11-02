package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/services"
	"gorm.io/gorm"
)

func GetUserProfileInfo(ctx *gin.Context, db *gorm.DB) {
	// Get user ID from JWT middleware
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, resterr.NewUnauthorizedError("Usuário não autenticado"))
		return
	}

	// Call service to get profile
	profile, err := services.GetUserProfile(userID.(string), db)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	// Return the profile data
	ctx.JSON(200, profile)
}
