package server

import (
	"log"
	"net/http"
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
	v1 := http.NewServeMux()

	router := http.NewServeMux()
	router.Handle("api/v1/", http.StripPrefix("api/v1", v1))

	log.Printf("The server is listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)
}
