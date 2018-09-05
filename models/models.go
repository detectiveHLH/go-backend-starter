package models

import (
	_ "database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-backend-starter/pkg/setting"
	"log"
)

var db *gorm.DB

func Setup() {
	/**
	连接数据库
	@param	type		类型
	@param	user		用户名
	@param	password	密码
	@param	host		地址
	@param	name		数据库名
	 */
	var err error
	db, err = gorm.Open(
		setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Println(err)
	}
}