package server

import (
	"forum/pkg/handler"
	"forum/repository"
)

func Server() {
	repository.CreateTable()
	handler.ListenServer()
}
