package handler

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./resources/html/addpost.html")
	if err != nil {
		log.Fatal(err)
	}
	r.ParseForm()
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
	title := r.FormValue("title_name")
	content := r.FormValue("content_text")
	if title != "" && content != "" {
		db, _ := sql.Open("sqlite3", "./forum.db")
		InsertPost(db, title, content)
	}
}

func InsertPost(db *sql.DB, title string, content string) *Database {
	statement, _ := db.Prepare("INSERT INTO post(title, content) values(?,?);")
	statement.Exec(title, content)
	// statement.Exec("nurmeden.02@gmail.com", "nurmeden", "vr3QcuFVQEDE8qz")
	return &Database{
		DB: db,
	}
}
