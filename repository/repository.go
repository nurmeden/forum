package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable() {
	database, _ := sql.Open("sqlite3", "./forum.db")
	NewTable(database, TableForComments)
	NewTable(database, TableForPosts)
	NewTable(database, TableForUsers)
	NewTable(database, TableForSession)
	NewTable(database, TableForLikes)
	NewTable(database, TableForLikesComment)
	defer database.Close()
}
