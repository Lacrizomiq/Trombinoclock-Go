package controllers

import (
	"net/http"
	"path/filepath"
	"trombinoclock/views"
)

func executeTemplate(w http.ResponseWriter, filename string, data interface{}) {
	tpl, err := views.Parse(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, data)
}

// Home is the controller for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("views/templates", "index.gohtml")
	executeTemplate(w, tplPath, nil)
}

func PromosHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("views/templates", "promos.gohtml")
	executeTemplate(w, tplPath, nil)
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("views/templates", "students.gohtml")
	executeTemplate(w, tplPath, nil)
}
