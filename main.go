package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/fvbock/endless"
)

func main() {
	fmt.Println("This works.")
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"code": 200,
			"message": "This works",
			"data": nil,
		})
	})
	r.Run()
}
