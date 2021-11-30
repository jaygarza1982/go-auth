package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	ginServer.GET("/", func(c *gin.Context) {
		status := "Status is OK"
		c.JSON(200, gin.H{"message": status})
	})

	ginServer.Run(":8080")
}
