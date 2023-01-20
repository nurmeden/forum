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
	Id          int
	Owner       string
	TitleName   string
	Description string
	Likes       string
	Dislikes    string
}

func Post(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./forum.db")
	rows, err := db.Query("SELECT * FROM post")
	fmt.Println(rows.Columns())
	var idDb int
	var ownerDb string
	var titleDb string
	var contentDb string
	var likesDb int
	var dislikesDb int
	var posts []post
	for rows.Next() {
		err = rows.Scan(&idDb, &titleDb, &ownerDb, &contentDb, &likesDb, &dislikesDb)
		if err != nil {
			log.Fatal(err)
		}
		itemInPosts := post{
			Id:          idDb,
			Owner:       ownerDb,
			TitleName:   titleDb,
			Description: contentDb,
			Likes:       "0",
			Dislikes:    "0",
		}
		posts = append(posts, itemInPosts)
	}
	tmpl, err := template.ParseFiles("./resources/html/post.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, posts)
	if err != nil {
		log.Fatal(err)
	}
}
