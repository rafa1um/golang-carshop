package controllers

import (
	"html/template"
	"net/http"

	"github.com/rafa1um/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	carros := models.GetAllCars()
	templates.ExecuteTemplate(w, "Index", carros)
}
