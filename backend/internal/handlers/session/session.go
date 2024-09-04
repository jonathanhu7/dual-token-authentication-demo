package session

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /", h.postSession)
}

func (h *Handler) postSession(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world"))
}
