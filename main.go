package main

import (
	"fmt"
	"go-backend-starter/models"
	"go-backend-starter/pkg/setting"
	"go-backend-starter/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

type Person struct {
	Name  string
	Phone string
}

// @title Golang Gin API
// @version 1.0
// @description Go Backend Starter
// @termsOfService https://github.com/detectiveHLH/go-backend-starter

// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main () {
	/**
	初始化系统设置
	 */
	setting.Setup()

	/**
	初始化模型层，连接数据库
	 */
	models.Setup()

	/**
	路由注入
	 */
	router := routers.InitRouter()

	/**
	启动服务器
	 */
	address := fmt.Sprintf("%s:%s", setting.ServerSetting.Ip, setting.ServerSetting.Port)
	server := endless.NewServer(address, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	// 处理服务器错误
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}