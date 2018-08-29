package main

import (
	"github.com/fvbock/endless"
	"log"
	"syscall"
	"github.com/detectiveHLH/backend/routers"
)

func main () {
	router := routers.InitRouter()

	server := endless.NewServer("localhost:4242", handler)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
		// save it somehow
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

	endless.ListenAndServe(":8000", router)
}


