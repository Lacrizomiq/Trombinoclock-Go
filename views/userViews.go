package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates *template.Template

// Parse initialize the template system by parsing all templates
func Parse() {
	templates = template.Must(template.ParseGlob(filepath.Join("views/templates", "*.gohtml")))
}

type Template struct {
	htmlTpl *template.Template
}

// NewTemplate returns a Template instance
func NewTemplate(name string) Template {
	t := templates.Lookup(name)
	if t == nil {
		panic("template not found: " + name)
	}
	return Template{htmlTpl: t}
}

// Execute executes the template with the given data.
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
