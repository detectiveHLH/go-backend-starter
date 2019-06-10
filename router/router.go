package router

import (
	_ "github.com/detectiveHLH/go-backend-starter/docs"
	"github.com/detectiveHLH/go-backend-starter/middleware/jwt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 第一版API
	apiVersionOne := router.Group("/api/v1/")
	apiVersionOne.Use(jwt.Jwt())
	// 定义鉴权接口
	router.GET("login", Login)
	apiVersionOne.GET("hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"code": 200,
			"message": "This works",
			"data": nil,
		})
	})
	return router
}
