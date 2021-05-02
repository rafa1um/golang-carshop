package main

import (
	"net/http"

	"github.com/rafa1um/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
