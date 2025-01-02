package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dixxe/dweb-personal-website/resources/templates"
	"github.com/dixxe/dweb-personal-website/service/repositories"
)

func GetShowBlog(w http.ResponseWriter, r *http.Request) {
	posts := repositories.Blog.GetAllValues()

	component := templates.ShowBlogPage(posts)
	component.Render(context.Background(), w)
}

func PostCreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	header := r.FormValue("header")
	content := r.FormValue("content")
	newPost := repositories.Post{Id: 0, Header: header, Content: content}
	repositories.Blog.InsertValue(newPost)
	fmt.Println("Created post")
}

func PostDeletePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		fmt.Println(err)
	}
	repositories.Blog.DeleteValueByID(id)
}
