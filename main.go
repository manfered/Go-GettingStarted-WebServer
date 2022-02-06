package main

import (
	"net/http"

	"github.com/manfered/Go-GettingStarted-WebServer/controllers"
)

func main() {
	// register web server routing
	controllers.RegisterControllers()

	// serving the webserver on localhost port 5000
	http.ListenAndServe(":5000", nil)
}
