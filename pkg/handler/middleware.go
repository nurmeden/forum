package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}

func CheckerSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id int
		var username string
		var sessions string
		database, _ := sql.Open("sqlite3", "./forum.db")
		rows, _ := database.Query("SELECT * FROM sessions")
		for rows.Next() {
			err := rows.Scan(&id, &username, &sessions)
			if err != nil {
				log.Print(err.Error())
			}
			fmt.Println(id, username, sessions)
		}
		next.ServeHTTP(w, r)
	})
}

func welcome(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			next.ServeHTTP(w, r)
		} else {
			sessionToken := c.Value

			userSession, exists := sessions[sessionToken]
			if !exists {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			fmt.Println(userSession)
			if userSession.isExpired() {
				delete(sessions, sessionToken)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}
