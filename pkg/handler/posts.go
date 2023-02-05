package handler

import (
	"database/sql"
	"forum/models"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func Posts(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./forum.db")
	rows, err := db.Query("SELECT * FROM post")
	var idDb int
	var ownerDb string
	var titleDb string
	var contentDb string
	var likesDb int
	var dislikesDb int
	var posts []models.Post
	for rows.Next() {
		err = rows.Scan(&idDb, &ownerDb, &titleDb, &contentDb, &likesDb, &dislikesDb)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		itemInPosts := models.Post{
			Id:          idDb,
			Owner:       ownerDb,
			TitleName:   titleDb,
			Description: contentDb,
			Likes:       "0",
			Dislikes:    "0",
		}
		posts = append(posts, itemInPosts)
	}
	rows.Close()
	tmpl, err := template.ParseFiles("./resources/html/posts.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, posts)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
