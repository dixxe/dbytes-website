package service

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var database_name = "blogs.db"

type Post struct {
	Id      int
	Header  string
	Content string
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite", database_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database was openned")

	return db
}

func initDb() {
	db := openDb()
	defer db.Close()

	_, err := db.Exec(`
	CREATE TABLE blogs(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		header TEXT,
		content TEXT
	);
	`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database was initiated.")
}

func GetAllPosts() []Post {
	db := openDb()
	defer db.Close()

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

func GetPost(id int) Post {
	db := openDb()
	defer db.Close()

	row := db.QueryRow("SELECT * from blogs where id = ?", id)
	p := Post{}
	err := row.Scan(&p.Id, &p.Header, &p.Content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Get one post")

	return p
}

func DeletePost(id int) {
	db := openDb()
	defer db.Close()

	_, err := db.Exec("DELETE from blogs where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Delete one post")

}

// Returning last inserted id
func CreatePost(header string, content string) int64 {
	db := openDb()
	defer db.Close()

	result, err := db.Exec("INSERT into blogs (header, content) values (?,?)",
		header, content)

	if err != nil {
		fmt.Println(err)
	}

	id, _ := result.LastInsertId()

	return id
}
