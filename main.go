package main

import (
	"fmt"

	"github.com/mohieey/UsersWebService/models"
)

func main() {
	u := models.User{ID: 23,
		FirstName: "Lionel",
		LastName:  "Messi",
	}

	fmt.Println(u)
}
