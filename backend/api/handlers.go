package api

import (
	"log"
	"usernamecheck/bloomfilter"
	"usernamecheck/postgres"
	"usernamecheck/redis"

	"github.com/gin-gonic/gin"
)

func CheckUsernameHandler(c *gin.Context) {
	username := c.PostForm("username")

	if username == "" {
		c.JSON(400, gin.H{
			"error": "Username is required",
		})
		return
	}

	if !bloomfilter.IsUsernameInBloom(username) {
		c.JSON(200, gin.H{"available": true, "message": "Username is available"})
		return
	}

	isCached, err := redis.IsUsernamePresentInCache(username)
	if err != nil {
		log.Println("Error checking username in cache:", err)
	} else if isCached {
		log.Println("Username found in cache:", username)
		c.JSON(200, gin.H{"available": false, "message": "Username is already taken"})
		return
	}

	available, err := postgres.CheckUsername(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if !available {
		log.Println("Username not available:", username)
		cacheErr := redis.StoreUsernameInCache(username)
		if cacheErr != nil {
			log.Println("Error storing username in cache:", cacheErr)
		}
		c.JSON(200, gin.H{"available": false, "message": "Username is already taken"})
		return
	}

	c.JSON(200, gin.H{"available": available, "message": "Username is available"})
}
