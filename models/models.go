package models

import (
	"../pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

//type Model struct {
//	ID         int `gorm:"primary_key" json:"id"`
//	CreatedOn  int `json:"created_on"`
//	ModifiedOn int `json:"modified_on"`
//	DeletedOn  int `json:"deleted_on"`
//}

func Setup() {
	/**
	连接数据库
	 */
	 fmt.Println(setting.DatabaseSetting)
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