package main

import (
	"net/http"

	"github.com/dixxe/dweb-personal-website/service/controllers"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
