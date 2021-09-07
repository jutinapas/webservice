package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// this is a method, a function that refered to the object ('this' keyword is not acceptable in Go)
// implicit implemented Handler interface
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// convention of the contructor function
func newUserController() *userController {
	// grab user id and return with an address of contructed controller object
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
