package router

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func SetupRouter() {
	// Pour rentre disponible les .css  et .js au html du front end
	http.Handle("/router/static/", http.StripPrefix("/router/static", http.FileServer(http.Dir("router/static"))))
	http.HandleFunc("/", routerHome) // il est mieux d'utiliser des http.mux
}

func routerHome(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles(fmt.Sprintf("./router/templates/%s", tmpl))

	if t == nil {
		log.Println("templates is not found")
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t.Execute(w, nil)
}
