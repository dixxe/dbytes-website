/*
Controller that only shows index template! Simple as it is!
*/

package controllers

import (
	"context"
	"net/http"

	"github.com/dixxe/dbytes-website/resources/templates"
)

func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	// Templates are components and this is basic way to render them.
	// I use this way across all code
	component := templates.IndexPage()
	component.Render(context.Background(), w)
}
