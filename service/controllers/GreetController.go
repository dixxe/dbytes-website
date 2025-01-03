/*
Simple controller that I wrote first for practice.
It serves as example for API routing,
*/

package controllers

import (
	"context"
	"net/http"

	"github.com/dixxe/dweb-personal-website/resources/templates"
	"github.com/go-chi/chi/v5"
)

func GetProcessGreet(w http.ResponseWriter, r *http.Request) {
	nameToGreet := chi.URLParam(r, "name")
	if nameToGreet == "" { // Placeholder for variable
		nameToGreet = "мир"
	}
	greeting := "Привет " + nameToGreet + "!"
	component := templates.GreetPage(greeting)
	component.Render(context.Background(), w)
}
