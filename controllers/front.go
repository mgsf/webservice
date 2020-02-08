package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

//RegisterControllers does all the hard work in routing
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
