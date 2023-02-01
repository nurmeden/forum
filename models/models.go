package models

type User struct {
	Id       int
	Email    string
	Username string
	Password string
}

type Post struct {
	Id          int
	Owner       string
	TitleName   string
	Description string
	Likes       string
	Dislikes    string
}

type Error struct {
	Message string
	Code    int
}

var Users []*User
