package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetIcon(w http.ResponseWriter, r *http.Request) {
	icon, err := os.Open("./resources/static/images/icon.png")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer icon.Close()

	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, icon)
}
