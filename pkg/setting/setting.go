package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

// 定义服务器配置结构体
type Server struct {
	Ip   string
	Port string
}
type Database struct {
	Url string
}

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var config *ini.File

func Setup() {
	var err error
	config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'config/app.ini': %v", err)
	}
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
