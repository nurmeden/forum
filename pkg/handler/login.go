package handler

import (
	"log"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./resources/html/login.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
