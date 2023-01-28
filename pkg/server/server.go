package server

import (
	"forum/pkg/handler"
	"forum/repository"
	"net/http"
)

func ServerRun() error {
	repository.CreateTable()
	err := http.ListenAndServe(":8080", handler.Routes())
	if err != nil {
		return err
	}
	return nil
}
