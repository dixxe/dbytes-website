package controllers

import (
	"context"
	"net/http"

	"github.com/dixxe/dweb-personal-website/resources/templates"
	"github.com/go-chi/chi/v5"
)

func GetProcessGreet(w http.ResponseWriter, r *http.Request) {
	nameToGreet := chi.URLParam(r, "name")
	if nameToGreet == "" {
		nameToGreet = "мир"
	}
	greeting := "Привет " + nameToGreet + "!"
	component := templates.GreetPage(greeting)
	component.Render(context.Background(), w)
}
