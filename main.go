package main

import (
	"fmt"

	"github.com/mgsf/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Amasha",
		LastName:  "Fernando",
	}

	fmt.Println(u)
}
