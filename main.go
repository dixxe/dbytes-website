package main

import (
	"net/http"

	"github.com/dixxe/learnbackend/service/controllers"
)

func main() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.ListenAndServe(":8080", nil)
}
