package handler

import (
	"log"
	"net/http"
)

func ListenServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home_Page)
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
