package main

import (
	"forum/pkg/server"
	"log"
)

func main() {
	err := server.ServerRun()
	if err != nil {
		log.Fatal(err)
	}
}
