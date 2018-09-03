package main

import (
	"./package/setting"
	"./routers"
	"fmt"
	"github.com/fvbock/endless"
	"gopkg.in/mgo.v2"
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

	/**
	连接数据库
	 */
	MOGODB_URI := "127.0.0.1:27017"
	//连接
	session, databaseErr := mgo.Dial(MOGODB_URI)
	//连接失败时终止
	if databaseErr != nil {
		panic(databaseErr)
	}
	//延迟关闭，释放资源
	defer session.Close()
	//设置模式
	session.SetMode(mgo.Monotonic, true)

	/**
	路由注入
	 */
	router := routers.InitRouter(session)

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
