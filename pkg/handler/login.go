package handler

import (
	"database/sql"
	"encoding/json"
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
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
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
				fmt.Println(inUser)
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
			tmpl, err := template.ParseFiles("./resources/html/index.html")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(userInPage)
			err = tmpl.Execute(w, userInPage)
			if err != nil {
				log.Fatal(err)
			}
			var creds Credentials

			err = json.NewDecoder(r.Body).Decode(&creds)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			expectedPassword, ok := users[creds.Username]

			if !ok || expectedPassword != creds.Password {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Create a new random session token
			// we use the "github.com/google/uuid" library to generate UUIDs
			sessionToken := uuid.NewString()
			expiresAt := time.Now().Add(120 * time.Second)

			// Set the token in the session map, along with the session information
			sessions[sessionToken] = session{
				username: creds.Username,
				expiry:   expiresAt,
			}

			// Finally, we set the client cookie for "session_token" as the session token we just generated
			// we also set an expiry time of 120 seconds
			http.SetCookie(w, &http.Cookie{
				Name:    "session_token",
				Value:   sessionToken,
				Expires: expiresAt,
			})
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
	}
}
