package main

import (
	"net/http"

	"github.com/ckaraboran/golang-http-helloworld/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
