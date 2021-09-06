package chapters

import "fmt"

func structs() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Fleck"
	fmt.Println(u)

	u2 := user{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Morgan",
	}
	fmt.Println(u2)
}
