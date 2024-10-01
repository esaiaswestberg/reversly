package main

import (
	"log"
	"reversly/internal/tcp/server"
)

func main() {
	log.Println("Welcome to Reversly!")

	err := server.ListenAndProxy(
		"0.0.0.0:8080",
		"127.0.0.1:1234",
	)
	if err != nil {
		log.Fatal(err)
	}
}
