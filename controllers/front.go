package controllers

import "net/http"

//RegisterControllers does all the hard work in routing
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
