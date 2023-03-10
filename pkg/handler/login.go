package handler

import (
	"database/sql"
	"fmt"
	"forum/models"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const COOKIE_NAME = "session_token"

var users = map[string]string{}

var sessions = map[string]session{}

// each session contains the username of the user and the time at which it expires
type session struct {
	Username string
	expiry   time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./resources/html/login.html")
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
		inUser := false
		database, _ := sql.Open("sqlite3", "./forum.db")
		rows, err := database.Query("SELECT * FROM users")
		defer database.Close()
		user := r.FormValue("uname")
		passwrd := r.FormValue("psw")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		var id int
		var email string
		var username string
		var password string
		for rows.Next() {
			err = rows.Scan(&id, &email, &username, &password)
			users[username] = password
			if user == username && passwrd == password {
				inUser = true
				break
			}
		}
		rows.Close()
		if inUser {
			userInPage := models.User{
				Id:       id,
				Email:    email,
				Username: username,
				Password: password,
			}
			fmt.Println(userInPage)
			expectedPassword, ok := users[username]

			if !ok || expectedPassword != password {
				ErrorHandler(w, http.StatusUnauthorized)
				return
			}

			sessionToken := uuid.NewString()
			expiresAt := time.Now().Add(120 * time.Second)

			sessions[sessionToken] = session{
				Username: username,
				expiry:   expiresAt,
			}
			database, err := sql.Open("sqlite3", "./forum.db")
			if err != nil {
				log.Fatal(err)
			}
			if err = database.Ping(); err != nil {
				log.Fatal(err)
			}
			defer database.Close()
			InsertSession(database, username, sessionToken)
			http.SetCookie(w, &http.Cookie{
				Name:    COOKIE_NAME,
				Value:   sessionToken,
				Expires: expiresAt,
				Path:    "/",
			})
			http.Redirect(w, r, "/", 302)
			return
		} else {
			http.Redirect(w, r, "/login", 302)
		}
		return
	}
	return
}

func InsertSession(db *sql.DB, user string, session string) *Database {
	statement, _ := db.Prepare("INSERT INTO sessions(user, session) values(?, ?);")
	_, err := statement.Exec(user, "bdfbdfbfdbdfb")
	if err != nil {
		log.Print("err", err)
	}
	return &Database{
		DB: db,
	}
}
