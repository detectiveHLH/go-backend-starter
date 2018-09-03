package user

import (
	"../../models"
	"fmt"
)

type User struct {
	Id int
	Name string
	Age  int
	State int
}

func (newUser *User) Add() error {
	fmt.Println("InService")
	user := map[string]interface{}{
		"name": newUser.Name,
		"age": newUser.Age,
		"state": newUser.State,
	}
	if err := models.AddUser(user); err != nil {
		return err
	}
	return nil
}

func (queryUser *User) Get() (*models.User, error) {
	user, err := models.GetUser(queryUser.Id)
	fmt.Println(queryUser.Id)
	if err != nil {
		fmt.Println("错误2")
		return nil, err
	}
	return user, nil
}