package handler

import (
	"database/sql"
	"fmt"
	"forum/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./resources/html/login.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		inUser := false
		database, _ := sql.Open("sqlite3", "./forum.db")
		rows, err := database.Query("SELECT * FROM users")

		user := r.FormValue("uname")
		passwrd := r.FormValue("psw")

		if err != nil {
			log.Fatal(err)
		}
		var id int
		var email string
		var username string
		var password string
		for rows.Next() {
			err = rows.Scan(&id, &email, &username, &password)
			if user == username && passwrd == password {
				inUser = true
				break
			}
		}
		if inUser {
			fmt.Println("ok")
			userInPage := models.User{
				Id:       id,
				Email:    email,
				Username: username,
				Password: password,
			}

			tmpl, err := template.ParseFiles("./resources/html/index.html")
			if err != nil {
				log.Fatal(err)
			}
			err = tmpl.Execute(w, userInPage)
			if err != nil {
				log.Fatal(err)
			}
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println(" return login url ")
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
