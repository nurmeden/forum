package handler

import (
	"database/sql"
	"log"
	"net/http"
)

func Like(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/like" {
			ErrorHandler(w, http.StatusNotFound)
			return
		}

		db, _ := sql.Open("sqlite3", "./forum.db")
		defer db.Close()
		AddLikeData(db)

	}
}

func AddLikeData(db *sql.DB) *Database {
	statement, _ := db.Prepare("DELETE FROM sessions;")
	_, err := statement.Exec()
	if err != nil {
		log.Print("err", err)
	}
	return &Database{
		DB: db,
	}
}
