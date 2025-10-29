package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/controllers/request"
	"github.com/rlevidev/taskmanager-go/src/models"
	"github.com/rlevidev/taskmanager-go/src/services"
	"gorm.io/gorm"
)

func CreateTask(ctx *gin.Context, db *gorm.DB) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(401, resterr.NewUnauthorizedError("Usuário não autenticado"))
		return
	}

	var taskRequest request.CreateTaskRequest

	if err := ctx.ShouldBindJSON(&taskRequest); err != nil {
		ctx.JSON(400, resterr.NewBadRequestError("Dados inválidos: " + err.Error()))
		return
	}

	// Criar nova tarefa
	taskDomain := models.NewTaskDomain(
		taskRequest.Title,
		taskRequest.Description,
		userID.(string),
	)

	// Salvar no banco
	createdTask, err := services.CreateTask(*taskDomain, db)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	// Retornar resposta
	ctx.JSON(201, gin.H{
		"message": "Tarefa criada com sucesso",
		"data":    createdTask,
	})
}
