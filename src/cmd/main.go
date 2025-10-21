package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rlevidev/taskmanager-go/src/config/database"
	"github.com/rlevidev/taskmanager-go/src/routes"
)

func main() {
	// Carregar arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	db, err := database.Init()
	if err != nil {
		log.Printf("❌ Erro ao conectar ao banco de dados: %v", err)
		log.Fatal("💀 Não foi possível conectar ao banco de dados. Verifique suas credenciais no arquivo .env")
	}
	log.Println("✅ Banco de dados inicializado com sucesso!")

	router := gin.Default()

	routes.SetupRoutes(router, db)
	log.Printf("🚀 Servidor iniciado na porta %s", os.Getenv("PORT"))

	router.Run()
}
