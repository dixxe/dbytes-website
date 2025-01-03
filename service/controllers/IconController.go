/*
Controller that adds icon for webpages.
It's displayed near the page <title>
*/

package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetIcon(w http.ResponseWriter, r *http.Request) {
	// This is relative path from main.go to image
	icon, err := os.Open("./resources/static/images/icon.png")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	// Don't forget to close files!
	defer icon.Close()

	// This code sets correct display settings, so browsers understand what I'm
	// goint to send them.
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, icon)
}
