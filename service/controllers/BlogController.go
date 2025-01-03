/*
Controller that handles everything d1xxe blog related features.
Under the hood it operates with repositories.Blog to get all information from
local database.

In a nutshell it's pretty simple.
*/
package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dixxe/dbytes-website/resources/templates"
	"github.com/dixxe/dbytes-website/service/repositories"
)

func GetShowBlog(w http.ResponseWriter, r *http.Request) {
	posts := repositories.Blog.GetAllValues()

	component := templates.ShowBlogPage(posts)
	component.Render(context.Background(), w)
}

func PostCreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // Populating form.
	header := r.FormValue("header") // To get value you need to specify name="header" in the form.
	content := r.FormValue("content")

	// Creating a new post with 0 Id, don't worry database handles id assignment
	// itself. And in the InsertValue() method I don't use Post.Id value
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
