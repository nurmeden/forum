package handler

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

func AddPost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./resources/html/addpost.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	case "POST":
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				ErrorHandler(w, http.StatusUnauthorized)
				return
			}
			ErrorHandler(w, http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		userSession, exists := sessions[sessionToken]
		if !exists {
			ErrorHandler(w, http.StatusUnauthorized)
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
			defer db.Close()
			InsertPost(db, userSession.Username, title, content)
		}
		http.Redirect(w, r, "/posts", 302)
	}
}

func InsertPost(db *sql.DB, owner string, title string, content string) *Database {
	statement, _ := db.Prepare("INSERT INTO post(owner, title, content) values(?, ?, ?);")

	statement.Exec(owner, title, content)
	return &Database{
		DB: db,
	}
}
