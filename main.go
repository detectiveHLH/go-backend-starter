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

	// 连接数据库
	//databaseUrl := "127.0.0.1:27017"
	//session, err :=

	// 启动服务器
	address := fmt.Sprintf("%s:%s", setting.ServerSetting.Ip, setting.ServerSetting.Port)
	server := endless.NewServer(address, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
