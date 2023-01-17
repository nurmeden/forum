package repository

import (
	"database/sql"
	"fmt"
)

func CreateTable() {
	database, _ := sql.Open("sqlite3", "./forum.db")

	NewTable(database, TableForComments)
	NewTable(database, TableForPosts)
	NewTable(database, TableForUsers)
	NewTable(database, TableForSession)
	NewTable(database, TableForLikes)
	NewTable(database, TableForLikesComment)
	rows, err := database.Query("SELECT * FROM post")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
}
