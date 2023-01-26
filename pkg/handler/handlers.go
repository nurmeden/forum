package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:

		if r.URL.Path != "/" {
			fmt.Println(http.StatusNotFound)
			return
		}

		// We can obtain the session token from the requests cookies, which come with every request
		tmpl, err := template.ParseFiles("./resources/html/index.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
