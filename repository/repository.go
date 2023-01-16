package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `
	CREATE TABLE IF NOT EXISTS post (
	id INTEGER NOT NULL PRIMARY KEY,
	time DATETIME NOT NULL,
	description TEXT,
	user VARCHAR(30),
);`

const like string = `
    CREATE TABLE IF NOT EXISTS like (
	id INTEGER NOT NULL PRIMARY KEY,
	time DATETIME NOT NULL,
	description TEXT,
);`

const comments string = `
	CREATE TABLE IF NOT EXISTS comments (
	id INTEGER NOT NULL PRIMARY KEY,
	time DATETIME NOT NULL,
	description TEXT,
	user VARCHAR(30),
);`

func CreateTable() {
	fmt.Println("os")
	db, _ := sql.Open("sqlite3", "./forum.db")
	fmt.Println("ok")

	NewTable(db, create)
	fmt.Println("no")
	NewTable(db, like)

	NewTable(db, comments)
	// for _, table := range tables {
	// 	NewTable(db, table)
	// }
	defer db.Close()
}
