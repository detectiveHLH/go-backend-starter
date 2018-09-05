package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Name string `json:"name"`
	Age  int `json:"age"`
	State int `json:"state"`
}

/**
添加用户
@param		data		数据
@return		error|nil	错误信息，如没有，则返回空
 */
func AddUser(data map[string]interface{}) error {
	newUser := User{
		Name: data["name"].(string),
		State: data["state"].(int),
		Age: data["age"].(int),
	}
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

/**
获取用户信息
@param		id			数据id
@returns	data, error	数据，错误信息
 */
func GetUser(id int) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}