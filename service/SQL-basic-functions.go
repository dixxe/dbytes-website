/*
Some SQL tools to reduce boilerplate.
Probably will be rewritten in future.
*/
package service

import (
	"database/sql"
	"fmt"
)

func OpenDb(database_name string) *sql.DB {
	db, err := sql.Open("sqlite", database_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database was openned")

	return db
}

// I don't know a way how to automate this process.
func InitDb(db *sql.DB) {
	//defer db.Close()

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
