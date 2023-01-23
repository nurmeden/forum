package handler

import (
	"database/sql"
	"fmt"
	"forum/models"
	"log"
	"net/http"
	"text/template"
)

type Home struct {
	Posts       []models.Post
	CurrentUser bool
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		if r.URL.Path != "/" {
			fmt.Println(http.StatusNotFound)
			return
		}

		Nick := false
		cook, err := r.Cookie(COOKIE_NAME)
		if err == nil {
			Nick = true
			sess := cook.Value
			DB, _ := sql.Open("sqlite3", "./forum.db")
			defer DB.Close()
			rows, _ := DB.Query("select * from sessions where session ='" + sess + "'")
			var id int
			var user string
			var session string
			for rows.Next() {
				rows.Scan(&id, &user, &session)
			}
			if user == "" {
				Nick = false
			}
		}
		fmt.Println(Nick)
		// We can obtain the session token from the requests cookies, which come with every request
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
