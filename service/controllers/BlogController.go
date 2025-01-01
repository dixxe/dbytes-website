package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dixxe/dweb-personal-website/resources/templates"
	"github.com/dixxe/dweb-personal-website/service"
)

func GetShowBlog(w http.ResponseWriter, r *http.Request) {
	posts := service.GetAllPosts()

	component := templates.ShowBlogPage(posts)
	component.Render(context.Background(), w)
}

func PostCreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	header := r.FormValue("header")
	content := r.FormValue("content")
	service.CreatePost(header, content)
	fmt.Println("Created post")
}

func PostDeletePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		fmt.Println(err)
	}
	service.DeletePost(id)
}
