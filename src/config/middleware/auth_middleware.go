package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Pegar token do header Authorization
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(401, resterr.NewUnauthorizedError("Token de acesso necessário"))
			ctx.Abort()
			return
		}

		// Verificar formato "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(401, resterr.NewUnauthorizedError("Formato de token inválido"))
			ctx.Abort()
			return
		}

		token := tokenParts[1]

		// Validar token
		claims, err := services.ValidateToken(token)
		if err != nil {
			ctx.JSON(err.Status, err)
			ctx.Abort()
			return
		}

		// Adicionar informações do usuário no contexto
		ctx.Set("user_id", claims.UserID)
		ctx.Set("user_email", claims.Email)
		ctx.Set("user_name", claims.Name)

		ctx.Next()
	}
}
