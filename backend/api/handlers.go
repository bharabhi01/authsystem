package api

import (
	"usernamecheck/postgres"

	"github.com/gin-gonic/gin"
)

func CheckUsernameHandler(c *gin.Context) {
	username := c.PostForm("username")

	ok, err := postgres.CheckUsername(username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if ok {
		c.JSON(200, gin.H{"available": true, "message": "Username is available"})
	} else {
		c.JSON(200, gin.H{"available": false, "message": "Username is already taken"})
	}
}
