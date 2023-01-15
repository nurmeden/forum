package repository

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "activities.db"

const create string = `
  CREATE TABLE IF NOT EXISTS activities (
  id INTEGER NOT NULL PRIMARY KEY,
  time DATETIME NOT NULL,
  description TEXT
  );`

type Activities struct {
	mu sync.Mutex
	db *sql.DB
}

func CreateTable() (*Activities, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	return &Activities{
		db: db,
	}, nil

}
