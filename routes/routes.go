package routes

import (
	"net/http"

	"github.com/rafa1um/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
