package controllers

import (
	"context"
	"net/http"

	"github.com/dixxe/dweb-personal-website/resources/templates"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.IndexPage()
	component.Render(context.Background(), w)
}
