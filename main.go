package main

import (
	"log"
	"net/http"
	"trombinoclock/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", controllers.HomeHandler)
	r.Get("/promos", controllers.PromosHandler)
	r.Get("/students", controllers.StudentsHandler)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
