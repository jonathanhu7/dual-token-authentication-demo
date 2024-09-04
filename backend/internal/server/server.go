package server

import (
	"log"
	"net/http"

	"github.com/jonathanhu7/dual-token-authentication-demo/backend/internal/handlers/token"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	tokenRouter := http.NewServeMux()
	tokenHandler := token.NewHandler()
	tokenHandler.RegisterRoutes(tokenRouter)

	v1 := http.NewServeMux()
	v1.Handle("/tokens/", http.StripPrefix("/tokens", tokenRouter))

	router := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	log.Printf("The server is listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}
