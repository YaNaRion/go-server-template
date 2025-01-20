package router

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var errorHTMLNotFound = errors.New("HTML FILE NOT FOUND")

const home = "home.html"

func SetupRouter() {
	// Pour rentre disponible les .css  et .js au html du front end
	http.Handle("/router/static/", http.StripPrefix("/router/static", http.FileServer(http.Dir("router/static"))))
	http.HandleFunc("/home", routerHome) // il est mieux d'utiliser des http.mux
}

func routerHome(w http.ResponseWriter, r *http.Request) {
	t, err := renderTemplate(home)
	if err != nil {
		if err == errorHTMLNotFound {
			log.Printf("%s HTML FILE NOT FOUND\n", home)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("An error occure while sending HTML file: %s \n", err)

	}
}

func renderTemplate(tmpl string) (*template.Template, error) {
	t, err := template.ParseFiles(fmt.Sprintf("./router/templates/%s", tmpl))

	if t == nil {
		return nil, errorHTMLNotFound
	}

	if err != nil {
		return nil, err
	}

	return t, nil
}
