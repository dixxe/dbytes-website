package service

/*
Special interface that I made to make system more reusable.
So if you need in future another db it should satisfy this interface.
Generic T represents field of database.
In this approach DB must be open everytime.
*/
type Repository[T any] interface {
	GetAllValues() []T
	GetValueByID(id int) T
	InsertValue(T) (id int)
	DeleteValueByID(id int)
}
