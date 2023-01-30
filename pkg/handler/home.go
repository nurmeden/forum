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
		if err != nil {
			tmpl, err := template.ParseFiles("./resources/html/index.html")
			if err != nil {
				fmt.Println("error in Parsing")
				log.Fatal(err)
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				fmt.Println("error in Executing")
				log.Fatal(err)
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
			http.Redirect(w, r, "/refresh", 302)
			return
		}
		//w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))

		// We can obtain the session token from the requests cookies, which come with every request
		tmpl, err := template.ParseFiles("./resources/html/index.html")
		if err != nil {
			fmt.Println("error in Parsing")
			log.Fatal(err)
		}
		fmt.Println(userSession.Username)
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
