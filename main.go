package main

import (
	"net/http"

	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/dixxe/dweb-personal-website/service/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(chiprometheus.NewMiddleware("service_name"))
	r.Handle("/metrics", promhttp.Handler())

	r.Get("/", controllers.IndexHandler)
	r.Get("/greet", controllers.GreetController)
	r.Get("/greet/{name}", controllers.GreetController)

	http.ListenAndServe(":8080", r)
}
