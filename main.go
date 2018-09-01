package main

import (
	"./package/setting"
	"./routers"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main () {
	// 引入路由
	router := routers.InitRouter()
	setting.Setup()

	// 启动服务器
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
