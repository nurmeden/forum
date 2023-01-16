package repository

import (
	"database/sql"
	"fmt"
)

type Database struct {
	DB *sql.DB
}

func NewTable(db *sql.DB, str string) *Database {
	fmt.Println(str)
	statement, _ := db.Prepare(str)
	statement.Exec()
	// fmt.Println(statement)
	// res, err := statement.Exec()
	// if err != nil {
	// 	fmt.Println("ORaz")
	// 	os.Exit(0)
	// }
	// fmt.Println(res)
	return &Database{
		DB: db,
	}
}
