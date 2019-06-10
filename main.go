package main

import (
	"fmt"
	"github.com/detectiveHLH/go-backend-starter/router"
	"github.com/detectiveHLH/go-backend-starter/setting"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

// @title Go Backend Starter API
// @version 1.0
// @description Go Backend Starter
// @termsOfService https://github.com/detectiveHLH/go-backend-starter
func main() {
	setting.Setup()

	r := router.InitRouter()
	/**
	启动服务器
	*/
	address := fmt.Sprintf("%s:%s", setting.ServerSetting.Ip, setting.ServerSetting.Port)
	server := endless.NewServer(address, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	// 处理服务器错误
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
