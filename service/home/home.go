package home

import (
	"fmt"
	"net/http"
)

type Handler struct {
	// dependencies
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("/", h.Greet)

	router.HandleFunc("/go", h.healthCheck)

}

func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>go server running!</h1>")
	if err != nil {
		return
	}
}

func (h *Handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Healthy!"))
	if err != nil {
		return
	}
}
