package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/services"
	"gorm.io/gorm"
)

func FinishTask(ctx *gin.Context, db *gorm.DB) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, resterr.NewUnauthorizedError("Usuário não autenticado"))
		return
	}

	taskID := ctx.Param("task_id")
	if taskID == "" {
		ctx.JSON(400, resterr.NewBadRequestError("ID da tarefa é obrigatório"))
		return
	}

	// Atualizar status para "completed"
	err := services.UpdateTaskStatus(taskID, userID.(string), "completed", db)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Tarefa finalizada com sucesso",
	})
}

func DoingTask(ctx *gin.Context, db *gorm.DB) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, resterr.NewUnauthorizedError("Usuário não autenticado"))
		return
	}

	taskID := ctx.Param("task_id")
	if taskID == "" {
		ctx.JSON(400, resterr.NewBadRequestError("ID da tarefa é obrigatório"))
		return
	}

	// Atualizar status para "in_progress"
	err := services.UpdateTaskStatus(taskID, userID.(string), "in_progress", db)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Status da tarefa atualizado para 'fazendo'",
	})
}

func DeleteTask(ctx *gin.Context, db *gorm.DB) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, resterr.NewUnauthorizedError("Usuário não autenticado"))
		return
	}

	taskID := ctx.Param("task_id")
	if taskID == "" {
		ctx.JSON(400, resterr.NewBadRequestError("ID da tarefa é obrigatório"))
		return
	}

	// Deletar tarefa
	err := services.DeleteTask(taskID, userID.(string), db)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Tarefa deletada com sucesso",
	})
}
