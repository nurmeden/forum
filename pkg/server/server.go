package server

import (
	"forum/pkg/handler"
	"forum/repository"
)

func ServerRun() {
	repository.CreateTable()
	handler.Routes()
}
