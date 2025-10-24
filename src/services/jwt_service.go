package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlevidev/taskmanager-go/src/config/resterr"
	"github.com/rlevidev/taskmanager-go/src/models"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Carregado do .env

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.UserDomain) (string, *resterr.RestErr) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token válido por 24 horas

	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", resterr.NewInternalServerError("Erro ao gerar token")
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, *resterr.RestErr) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, resterr.NewUnauthorizedError("Token inválido")
	}

	if !token.Valid {
		return nil, resterr.NewUnauthorizedError("Token inválido")
	}

	return claims, nil
}
