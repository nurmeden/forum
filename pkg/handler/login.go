package handler

import (
	"database/sql"
	"fmt"
	"forum/models"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"text/template"
	"time"
)

var sessions = map[string]session{}

type session struct {
	username string
	expiry   time.Time
}

var users = map[string]string{
	"user1":    "password1",
	"nurmeden": "dulat2002",
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("./resources/html/login.html")
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		inUser := false
		database, _ := sql.Open("sqlite3", "./forum.db")
		rows, err := database.Query("SELECT * FROM users")

		user := r.FormValue("uname")
		passwrd := r.FormValue("psw")
		if err != nil {
			log.Fatal(err)
		}
		var id int
		var email string
		var username string
		var password string
		for rows.Next() {
			err = rows.Scan(&id, &email, &username, &password)
			if user == username && passwrd == password {
				inUser = true
				break
			}
		}
		if inUser {
			userInPage := models.User{
				Id:       id,
				Email:    email,
				Username: username,
				Password: password,
			}

			//var creds Credentials

			//err = json.NewDecoder(r.Body).Decode(&creds)
			//if err != nil {
			//	w.WriteHeader(http.StatusBadRequest)
			//	return
			//}

			expectedPassword, ok := users[userInPage.Username]
			fmt.Println(expectedPassword)
			if !ok || expectedPassword != userInPage.Password {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Create a new random session token
			// we use the "github.com/google/uuid" library to generate UUIDs
			sessionToken := uuid.NewString()
			expiresAt := time.Now().Add(120 * time.Second)
			fmt.Println(sessionToken)
			// Set the token in the session map, along with the session information
			sessions[sessionToken] = session{
				username: userInPage.Username,
				expiry:   expiresAt,
			}

			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   sessionToken,
				Expires: expiresAt,
			})
			tmpl, err := template.ParseFiles("./resources/html/index.html")
			if err != nil {
				log.Fatal(err)
			}
			err = tmpl.Execute(w, userInPage)
			if err != nil {
				log.Fatal(err)
			}
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			fmt.Println(" return login url ")
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
