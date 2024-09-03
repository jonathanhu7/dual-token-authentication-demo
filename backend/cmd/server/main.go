package main

import (
	"github.com/jonathanhu7/dual-token-authentication-demo/backend/internal/server"
	"log"
)

func main() {
	srv := server.NewServer(":8080")
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
