package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./resources/html/addpost.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	case "POST":
		c, err := r.Cookie("session_token")

		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("Status Unauthorized")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		userSession, exists := sessions[sessionToken]
		if !exists {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if userSession.isExpired() {
			delete(sessions, sessionToken)
			http.Redirect(w, r, "/refresh", 302)
			return
		}
		title := r.FormValue("title_name")
		content := r.FormValue("content_text")
		if title != "" && content != "" {
			db, _ := sql.Open("sqlite3", "./forum.db")
			InsertPost(db, userSession.Username, title, content)
		}
		http.Redirect(w, r, "/posts", 302)
	}
}

func InsertPost(db *sql.DB, owner string, title string, content string) *Database {
	fmt.Println("in processing db")
	statement, _ := db.Prepare("INSERT INTO post(owner, title, content) values(?, ?, ?);")
	statement.Exec(owner, title, content)
	// statement.Exec("nurmeden.02@gmail.com", "nurmeden", "vr3QcuFVQEDE8qz")
	return &Database{
		DB: db,
	}
}
