package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rlevidev/taskmanager-go/src/config/middleware"
	"github.com/rlevidev/taskmanager-go/src/controllers"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Grupo de rotas públicas (sem autenticação)
	public := r.Group("/api/v1")
	{
		// Rotas de usuário
		public.POST("/users/register", func(ctx *gin.Context) {
			controllers.CreateUser(ctx, db)
		})
		public.POST("/users/login", func(ctx *gin.Context) {
			controllers.LoginUser(ctx, db)
		})
	}

	// Grupo de rotas protegidas (com autenticação)
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		// Rotas que precisam de autenticação podem ser adicionadas aqui
		// Exemplo: protected.GET("/users/profile", controllers.GetUserProfile)
	}
}
