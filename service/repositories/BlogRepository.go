/*
	Managing blogs database with a little bit awful execution.
	This is can be considered as any-service database example.
	In this realisation database is *always* stays open after xxxService defintion.
	Because if something close it the struct will point to nothing and
	obviously this is bad.
*/

package repositories

import (
	"database/sql"
	"fmt"

	"github.com/dixxe/dweb-personal-website/service"
	_ "modernc.org/sqlite"
)

var database_name = "blogs.db" // You can change it if you want.

// Defining it to use it later via service.Blog
// Why here? Because it's best place for any-blog related!
var Blog blogService = blogService{Database: service.OpenDb(database_name)}

// Post structure for database field.
type Post struct {
	Id      int
	Header  string
	Content string
}

// This struct implemets Repository[Post]
type blogService struct {
	Database *sql.DB
}

func (blogs blogService) GetAllValues() []Post {
	db := blogs.Database
	//defer db.Close()

	rows, err := db.Query("SELECT * from blogs")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	posts := []Post{}

	for rows.Next() {
		p := Post{}
		err := rows.Scan(&p.Id, &p.Header, &p.Content)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, p)
	}
	fmt.Println("Readed all posts")

	return posts
}

func (blogs blogService) GetValueByID(id int) Post {
	db := blogs.Database
	//defer db.Close()

	row := db.QueryRow("SELECT * from blogs where id = ?", id)
	p := Post{}
	err := row.Scan(&p.Id, &p.Header, &p.Content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Get one post")

	return p
}

func (blogs blogService) DeleteValueByID(id int) {
	db := blogs.Database
	//defer db.Close()

	_, err := db.Exec("DELETE from blogs where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Delete one post")

}

// Returning last inserted id
func (blogs blogService) InsertValue(postToInsert Post) int {
	db := blogs.Database
	//defer db.Close()

	result, err := db.Exec("INSERT into blogs (header, content) values (?,?)",
		postToInsert.Header, postToInsert.Content)

	if err != nil {
		fmt.Println(err)
	}

	id, _ := result.LastInsertId()

	return int(id)
}
