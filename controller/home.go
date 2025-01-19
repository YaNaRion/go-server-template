package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const ContentTypeJSON = "application/json"

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func SetUpController() {
	log.Println("SET UP CONTROLLER")
	http.HandleFunc("/echo", handleEcho)
}

func handleEcho(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hello from handlerECHO")
	fmt.Println(r.Body)

	p := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	jsonData, _ := json.Marshal(p)

	w.Header().Set("Contend-Type", ContentTypeJSON)

	w.Write(jsonData)
}
