package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

// add method to a struct
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller"))
}

// constructor
func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
