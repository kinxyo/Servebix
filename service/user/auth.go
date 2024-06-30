package user

import "net/http"

type Handler struct {
	//	dependencies
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", handler.Login)
	router.HandleFunc("POST /register", handler.Register)
}

func (handler *Handler) Login(writer http.ResponseWriter, request *http.Request) {

}

func (handler *Handler) Register(writer http.ResponseWriter, request *http.Request) {

}
