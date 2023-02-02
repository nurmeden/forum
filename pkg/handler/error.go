package handler

import (
	"forum/models"
	"log"
	"net/http"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, code int) {
	ErrorInPage := models.Error{
		Message: http.StatusText(code),
		Code:    code,
	}
	tmpl, err := template.ParseFiles("./resources/html/error.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, ErrorInPage)
	if err != nil {
		log.Fatal(err)
	}
}
