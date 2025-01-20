package router

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var errorHTMLnotFound = errors.New("HTML FILE NOT FOUND")

const home = "home.html"

func SetupRouter() {
	// Pour rentre disponible les .css  et .js au html du front end
	http.Handle("/router/static/", http.StripPrefix("/router/static", http.FileServer(http.Dir("router/static"))))
	http.HandleFunc("/home", routerHome) // il est mieux d'utiliser des http.mux
}

func routerHome(w http.ResponseWriter, r *http.Request) {
	t, err := renderTemplate(home)
	if err != nil {
		if err == errorHTMLnotFound {
			log.Println(fmt.Sprintf("%s HTML FILE NOT FOUND", home))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	t.Execute(w, nil)
}

func renderTemplate(tmpl string) (*template.Template, error) {
	t, err := template.ParseFiles(fmt.Sprintf("./router/templates/%s", tmpl))

	if t == nil {
		return nil, errorHTMLnotFound
	}

	if err != nil {
		return nil, err
	}

	return t, nil
}
