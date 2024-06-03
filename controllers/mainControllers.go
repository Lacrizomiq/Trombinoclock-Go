package controllers

import (
	"net/http"
	"trombinoclock/views"
)

// Home is the controller for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	views.Parse()
	tmpl := views.NewTemplate("index")
	tmpl.Execute(w, nil)
}

func PromosHandler(w http.ResponseWriter, r *http.Request) {
	views.Parse()
	tmpl := views.NewTemplate("promos")
	tmpl.Execute(w, nil)
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	views.Parse()
	tmpl := views.NewTemplate("students")
	tmpl.Execute(w, nil)
}
