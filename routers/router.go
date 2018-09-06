package routers

import (
	"github.com/gin-gonic/gin"
	"go-backend-starter/middleware/jwt"
	"go-backend-starter/routers/v1"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	// 定义鉴权接口
	router.GET("/auth", GetAuth)

	// 第一版API
	apiVersionOne := router.Group("/api/v1/")
	apiVersionOne.Use(jwt.JWT())
	apiVersionOne.GET("getUser", v1.GetUser)
	apiVersionOne.POST("addUser", v1.AddUser)

	return router
}