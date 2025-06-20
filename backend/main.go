package main

import (
	"log"
	"usernamecheck/api"
	"usernamecheck/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := postgres.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer postgres.CloseDB()

	router := gin.Default()
	api.SetupRoutes(router)
	router.Run(":8080")
}
