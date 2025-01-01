package main

import (
	"net/http"

	"github.com/dixxe/dweb-personal-website/service/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", controllers.IndexHandler)
	r.Get("/greet", controllers.GreetController)
	r.Get("/greet/{name}", controllers.GreetController)
	r.Get("/blog", controllers.BlogController)
	http.ListenAndServe(":8080", r)
}
