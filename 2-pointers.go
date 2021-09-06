package main

import "fmt"

func pointers() {
	fmt.Print("\n---\n\n")
	var firstName *string = new(string)
	*firstName = "Arthur"

	fmt.Println(firstName)
	fmt.Println(*firstName)

	lastName := "Fleck"
	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Morgan"
	fmt.Println(ptr, *ptr)

	fmt.Print("\n---\n\n")
}
