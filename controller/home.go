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
	http.HandleFunc("/data", handleEcho)
}

func handleEcho(w http.ResponseWriter, r *http.Request) {

	log.Println(fmt.Sprintf("ECHO REQUEST FROM: %s", r.RemoteAddr))

	p := buildPerson()

	jsonData, err := json.Marshal(p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	writeResponse(w, jsonData)

}

func buildPerson() Person {
	return Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}
}

func writeResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Contend-Type", ContentTypeJSON)
	w.Write(data)
}
