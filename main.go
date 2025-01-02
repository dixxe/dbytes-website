package main

/*
	Main file with routes and nothing else.
*/

import (
	"net/http"

	"github.com/dixxe/dweb-personal-website/service/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.GetIndexHandler)

	r.Get("/greet", controllers.GetProcessGreet)
	r.Get("/greet/{name}", controllers.GetProcessGreet)

	r.Get("/blog", controllers.GetShowBlog)
	r.Post("/post", controllers.PostCreatePost)
	r.Post("/post/delete", controllers.PostDeletePost)

	r.Post("/admin/login", controllers.PostAdminLogin)
	r.Get("/admin", controllers.GetAdminLogin)
	http.ListenAndServe(":8080", r)
}
