/*
Managing blogRepo database with a little bit awful execution.
This can be considered as any-repository database example.
In this realisation database is *always* stays open after xxxRepository intitialization.
Because if something close it the struct will point to nothing and obviously this is bad.
*/

package repositories

import (
	"database/sql"
	"fmt"

	"github.com/dixxe/dbytes-website/service"
	_ "modernc.org/sqlite"
)

var database_name = "blogs.db" // You can change it if you want.

// Defining it to use it later via repository.Blog
// Why here? Because it's best place for anything blog related!
var Blog blogRepository = blogRepository{Database: service.OpenDb(database_name)}

// Post structure for database field.
type Post struct {
	Id      int
	Header  string
	Content string
}

// This struct implemets Repository[Post]
type blogRepository struct {
	Database *sql.DB
}

func (blogRepo blogRepository) GetAllValues() []Post {
	db := blogRepo.Database
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

func (blogRepo blogRepository) GetValueByID(id int) Post {
	db := blogRepo.Database
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

func (blogRepo blogRepository) DeleteValueByID(id int) {
	db := blogRepo.Database
	//defer db.Close()

	_, err := db.Exec("DELETE from blogs where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Delete one post")

}

// Returning last inserted id
func (blogRepo blogRepository) InsertValue(postToInsert Post) int {
	db := blogRepo.Database
	//defer db.Close()

	result, err := db.Exec("INSERT into blogs (header, content) values (?,?)",
		postToInsert.Header, postToInsert.Content)

	if err != nil {
		fmt.Println(err)
	}

	id, _ := result.LastInsertId()

	return int(id)
}

func (blogRepo blogRepository) ExecSpecific(SQL_command string) sql.Result {
	result, err := blogRepo.Database.Exec(SQL_command)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func (blogRepo blogRepository) QuerySpecific(SQL_command string) *sql.Rows {
	result, err := blogRepo.Database.Query(SQL_command)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
