package main

import (
	"log"
	"usernamecheck/api"
	"usernamecheck/postgres"
	"usernamecheck/redis"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Redis
	if err := redis.InitRedis(); err != nil {
		log.Println("Redis initialization failed, continuing without caching")
	} else {
		log.Println("Redis initialized successfully")
		defer redis.CloseRedis()
	}

	// Initialize Postgres database
	if err := postgres.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	} else {
		log.Println("Postgres database initialized successfully")
		defer postgres.CloseDB()
	}

	router := gin.Default()
	api.SetupRoutes(router)
	router.Run(":8080")
}
