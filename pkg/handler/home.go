package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	//var user models.User
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/" {
			fmt.Println(http.StatusNotFound)
			return
		}

		// We can obtain the session token from the requests cookies, which come with every request
		tmpl, err := template.ParseFiles("./resources/html/index.html")
		if err != nil {
			fmt.Println("error in Parsing")
			log.Fatal(err)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println("error in Executing")
			log.Fatal(err)
		}
	case http.MethodPost:
		fmt.Println("method post")
		tmp, err := template.ParseFiles("./resources/html/index.html")
		err = tmp.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
