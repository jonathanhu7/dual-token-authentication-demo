package server

import (
	"log"
	"net/http"

	"github.com/jonathanhu7/dual-token-authentication-demo/backend/internal/handlers/session"
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
	sessionRouter := http.NewServeMux()
	sessionHandler := session.NewHandler()
	sessionHandler.RegisterRoutes(sessionRouter)

	v1 := http.NewServeMux()
	v1.Handle("/sessions/", http.StripPrefix("/sessions", sessionRouter))

	router := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	log.Printf("The server is listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}
