/*
Special interface that I made to make system more reusable.
So if you need in future another db it should satisfy this interface.
Generic T represents field of database.
In this approach DB must be open everytime.

Why interface? For example you create a function that executes something on database
and you need to execute this command automaticly for any DB that contains service.Post passed.
So you will define it like this
func exec_fix_for_db(repo service.Repository[Post]) {
	repo.ExecSpecific("Very important fix!!")
}
*/

package service

import "database/sql"

type Repository[T any] interface {
	GetAllValues() []T
	GetValueByID(id int) T
	InsertValue(T) (id int)
	DeleteValueByID(id int)
	ExecSpecific(SQL_command string) sql.Result
	QuerySpecific(SQL_command string) *sql.Rows
}
