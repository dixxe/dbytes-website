package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/dixxe/dweb-personal-website/resources/templates"
	"github.com/joho/godotenv"
)

func GetAdminPanel(w http.ResponseWriter, r *http.Request) {
	component := templates.AdminPanelPage()
	component.Render(context.Background(), w)
}

func GetAdminLogin(w http.ResponseWriter, r *http.Request) {
	component := templates.LoginPage()
	component.Render(context.Background(), w)
}

func PostAdminLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := r.FormValue("login")
	password := r.FormValue("password")

	if err := godotenv.Load(); err != nil {
		panic("No .env file located.")
	}

	admin_login, _ := os.LookupEnv("LOGIN")
	admin_password, _ := os.LookupEnv("PASSWORD")
	if login == admin_login && password == admin_password {
		GetAdminPanel(w, r)
	}
}
