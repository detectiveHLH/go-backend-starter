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