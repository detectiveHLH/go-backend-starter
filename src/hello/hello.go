package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/fvbock/endless"
)

func main() {
	fmt.Println("System loading...")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
