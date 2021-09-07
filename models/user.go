package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	// slice of users will make more efficiency for managing between users
	users []*User

	// at package level definition, no need implicit initialization for just int
	nextID = 1
)
