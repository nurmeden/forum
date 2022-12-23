package main

import (
	"forum/pkg/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home_Page)
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
