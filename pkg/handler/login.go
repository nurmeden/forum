package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"text/template"
	"time"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
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

			sessionToken := uuid.NewString()
			expiresAt := time.Now().Add(120 * time.Second)

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
