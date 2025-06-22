package main

import (
	"log"
	"usernamecheck/api"
	"usernamecheck/bloomfilter"
	"usernamecheck/postgres"
	"usernamecheck/redis"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Bloom Filter
	bloomfilter.InitBloomFilter(1000000, 0.01)

	bloomfilter.AddUsernameToBloom("testuser")
	bloomfilter.AddUsernameToBloom("john_doe")
	bloomfilter.AddUsernameToBloom("popular_user")
	bloomfilter.AddUsernameToBloom("jane_smith")
	bloomfilter.AddUsernameToBloom("admin")
	bloomfilter.AddUsernameToBloom("testuser1")
	log.Println("Bloom filter populated with initial usernames")

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
