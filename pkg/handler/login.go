package handler

import (
	"database/sql"
	"fmt"
	"forum/models"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
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
				fmt.Println(inUser)
				break
			}
		}
		if inUser {
			userInPage := models.User{
				Id:       id,
				Email:    email,
				Username: username,
				Password: password,
			}
			http.Redirect(w, r, "/", http.StatusFound)
			tmpl, err := template.ParseFiles("./resources/html/index.html")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(userInPage.Username)
			err = tmpl.Execute(w, userInPage)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
