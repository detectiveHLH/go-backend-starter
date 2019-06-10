package main

import (
	"github.com/detectiveHLH/go-backend-starter/router"
)

// @title Golang Gin API
// @version 1.0
// @description Go Backend Starter
// @termsOfService https://github.com/detectiveHLH/go-backend-starter
func main() {
	r := router.InitRouter()
	r.Run()
}
