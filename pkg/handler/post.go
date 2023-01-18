package handler

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"text/template"
)

type post struct {
	id          int
	owner       string
	title       string
	description string
	likes       string
	dislikes    string
}

func Post(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./resources/html/post.html")
	if err != nil {
		log.Fatal(err)
	}
	db, _ := sql.Open("sqlite3", "./forum.db")
	rows, err := db.Query("SELECT * FROM post")
	var id int
	var owner string
	var title string
	var content string
	var likes int
	var dislikes int
	for rows.Next() {
		err = rows.Scan(&id, &title, &owner, &content, &likes, &dislikes)
		fmt.Println(id, title, owner, content, likes, dislikes)
	}

	defer db.Close()
	err = tmpl.Execute(w, post{})
	if err != nil {
		log.Fatal(err)
	}
}
