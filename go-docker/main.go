package main

import (
	"go-docker/test"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helloWorld",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": test.CreateMessage(),
		})
	})
	r.Run(":3000")
}
