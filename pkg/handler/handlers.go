package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				fmt.Println("401")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		sessionToken := c.Value

		// We then get the session from our session map
		userSession, exists := sessions[sessionToken]
		if !exists {
			// If the session token is not present in session map, return an unauthorized error
			fmt.Println("401 - 2")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the session is present, but has expired, we can delete the session, and return
		// an unauthorized status
		if userSession.isExpired() {
			delete(sessions, sessionToken)
			fmt.Println("401 - 3")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// If the session is valid, return the welcome message to the user
		w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
		tmpl, err := template.ParseFiles("./resources/html/index.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
