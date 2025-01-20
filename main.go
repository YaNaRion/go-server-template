package main

import (
	"controller"
	"log"
	"net/http"
	"router"
)

const port = ":3000"

func main() {
	log.Println("Setting Router")
	router.SetupRouter()

	log.Println("Setting Controller")
	controller.SetUpController()

	log.Println("Started on Port: localhost:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
