package main

import (
	"log"
	"net/http"
	"trombinoclock/controllers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", controllers.HomeHandler)
	r.Get("/promos", controllers.PromosHandler)
	r.Get("/students", controllers.StudentsHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)
}
