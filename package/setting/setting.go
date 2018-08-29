package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

type Server struct {
	Port     int
}
var ServerSetting = &Server{}

var config *ini.File

func Setup() {
	var err error
	config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'config/app.ini': %v", err)
	}
	mapTo("server", ServerSetting)
}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
