package handler

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		if r.URL.Path != "/" {
			ErrorHandler(w, http.StatusNotFound)
			return
		}

		c, err := r.Cookie("session_token")
		if err != nil {
			tmpl, err := template.ParseFiles("./resources/html/index.html")
			if err != nil {
				ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			return
		}
		sessionToken := c.Value
		userSession, _ := sessions[sessionToken]
		// if !exists {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		if userSession.isExpired() {
			delete(sessions, sessionToken)
			db, _ := sql.Open("sqlite3", "./forum.db")
			defer db.Close()
			DeleteSession(db)
			http.Redirect(w, r, "/refresh", 302)
			return
		}
		// w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))

		// We can obtain the session token from the requests cookies, which come with every request
		tmpl, err := template.ParseFiles("./resources/html/index.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, userSession)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		tmp, err := template.ParseFiles("./resources/html/index.html")
		err = tmp.Execute(w, nil)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}
}

func DeleteSession(db *sql.DB) *Database {
	statement, _ := db.Prepare("DELETE FROM sessions;")
	_, err := statement.Exec()
	if err != nil {
		log.Print("err", err)
	}
	return &Database{
		DB: db,
	}
}
