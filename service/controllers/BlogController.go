package controllers

import (
	"context"
	"net/http"

	"github.com/dixxe/dweb-personal-website/resources/templates"
)

func BlogController(w http.ResponseWriter, r *http.Request) {
	component := templates.WIPPage()
	component.Render(context.Background(), w)
}
