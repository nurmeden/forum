package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./resources/html/signUp.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("psw")
		Repeat_password := r.FormValue("psw-repeat")
		if password == Repeat_password {
			db, _ := sql.Open("sqlite3", "./forum.db")
			defer db.Close()
			InsertData(db, email, username, password)
			fmt.Println(username)
		}
	case http.MethodPost:
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

func InsertData(db *sql.DB, email string, username string, password string) *Database {
	statement, _ := db.Prepare("INSERT INTO users(email, username, password) values(?,?,?);")
	statement.Exec(email, username, password)
	return &Database{
		DB: db,
	}
}
