package handler

import (
	"log"
	"net/http"
)

func ListenServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home_Page)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/signUp", SignUp)
	mux.HandleFunc("/Addpost", PostPage)
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	fileServer := http.FileServer(http.Dir("./resources/"))
	mux.Handle("/resources/", http.StripPrefix("/resources/", fileServer))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
