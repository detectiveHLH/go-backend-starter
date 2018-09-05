package app

import (
	"github.com/gin-gonic/gin"
	"go-backend-starter/pkg/e"
)

type Gin struct {
	C *gin.Context
}

/**
统一后端返回格式
@param	httpCode	http状态码
@param	errCode		错误码
@param	data		返回数据
 */
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}
