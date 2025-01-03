/*
This controller handles admin stuff and allows to redact databases.

It has some noodles in it but this is how it works.
1. Users enters on "/admin" and this controller handles GetAdminLogin()
2. Than frontend passes a form with login and password to PostAdminLogin()
3. In PostAdminLogin() backend checks that login and password are equal to
	ones that defined in .env file (without it code will panic! I will rewrite this)
4. If everything correct controller goes to GetAdminPanel() and shows admin-panel
	template

Currently page is vulnerable to bruteforce.
I will try to implement JWT in future. -d1xxe
*/

package controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/dixxe/dbytes-website/resources/templates"
	"github.com/dixxe/dbytes-website/service/repositories"
	"github.com/joho/godotenv"
)

func GetAdminPanel(w http.ResponseWriter, r *http.Request) {
	posts := repositories.Blog.GetAllValues()
	component := templates.AdminPanelPage(posts)
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
		// I better disable admin panel at all if .env not located.
		panic("No .env file located.")
	}

	admin_login, _ := os.LookupEnv("LOGIN")
	admin_password, _ := os.LookupEnv("PASSWORD")
	if login == admin_login && password == admin_password {
		GetAdminPanel(w, r)
	}
}
