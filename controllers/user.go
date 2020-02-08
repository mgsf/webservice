package controllers

import (
	"net/http"
	"regexp"
)

//userController will handle two types of resource requests, users collection & a single user
type userController struct {
	userIDpattern *regexp.Regexp
}

func (us userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller"))
}

//newUserController - func that wroks as the constructor
func newUserController() *userController {
	return &userController{
		userIDpattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
