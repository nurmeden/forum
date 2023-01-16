package models

type User struct {
	id             int    `json: "id" db : "id"`
	email          string `json: "email" db : "email"`
	Username       string `json:"username" db:"username"`
	HashedPassword string `json:"hashed_password" db:"hashed_password"`
}
