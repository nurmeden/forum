package repository

import (
	"database/sql"
)

type Database struct {
	DB *sql.DB
}

func NewTable(db *sql.DB, str string) *Database {
	statement, _ := db.Prepare(str)

	if str == TableForUsers {
		statement, _ = db.Prepare("INSERT INTO users(email, username, password) values('nurmeden.02@gmail.com','nurmeden','dulat2002')")
	}
	statement.Exec()
	return &Database{
		DB: db,
	}
}
