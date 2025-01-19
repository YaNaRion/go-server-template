package main

import (
	"controller"
	"log"
	"net/http"
	"router"
)

const port = ":3000"

func main() {
	log.Println("Hello world")

	router.SetupRouter()
	controller.SetUpController()

	log.Println("Started on Port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
