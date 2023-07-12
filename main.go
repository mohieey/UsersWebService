package main

import (
	"net/http"

	"github.com/mohieey/UsersWebService/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
