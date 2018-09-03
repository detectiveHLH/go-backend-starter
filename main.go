package main

import (
	"./models"
	"./pkg/setting"
	"./routers"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

type Person struct {
	Name  string
	Phone string
}

func main () {
	/**
	初始化系统设置
	 */
	setting.Setup()
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