package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	switch r.Method {
	case http.MethodGet:

		if r.URL.Path != "/" {
			fmt.Println(http.StatusNotFound)
			return
		}

		c, err := r.Cookie("session_token")
		fmt.Println(c)
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
		fmt.Println(userSession)
		if userSession.isExpired() {
			delete(sessions, sessionToken)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		//w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))

		// We can obtain the session token from the requests cookies, which come with every request
		tmpl, err := template.ParseFiles("./resources/html/index.html")
		if err != nil {
			fmt.Println("error in Parsing")
			log.Fatal(err)
		}

		err = tmpl.Execute(w, userSession)
		if err != nil {
			fmt.Println("error in Executing")
			log.Fatal(err)
		}
	case http.MethodPost:
		fmt.Println("method post")
		tmp, err := template.ParseFiles("./resources/html/index.html")
		err = tmp.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
