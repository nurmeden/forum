package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Home_Page(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./resources/html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok home_page ")
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
