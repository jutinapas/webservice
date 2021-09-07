package main

import (
	"fmt"

	"github.com/jutinapas/webservices/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Fleck",
	}
	fmt.Println(u)
}
