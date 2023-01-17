package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./resources/html/signUp.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("psw")
	Repeat_password := r.FormValue("psw-repeat")
	if password == Repeat_password {
		db, _ := sql.Open("sqlite3", "./forum.db")
		InsertData(db, email, username, password)
		fmt.Println(username)
	}
}

func InsertData(db *sql.DB, email string, username string, password string) *Database {
	statement, _ := db.Prepare("INSERT INTO users(email, username, password) values(?,?,?);")
	statement.Exec(email, username, password)
	// statement.Exec("nurmeden.02@gmail.com", "nurmeden", "vr3QcuFVQEDE8qz")
	return &Database{
		DB: db,
	}
}
