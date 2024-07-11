package doctor

import "net/http"

type Routes interface {
	handleLogin(http.ResponseWriter, *http.Request)
	handleSignup(http.ResponseWriter, *http.Request)
	handleProfileUpdate(http.ResponseWriter, *http.Request)
}

// METHODS
func (handler *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) handleProfileUpdate(w http.ResponseWriter, r *http.Request) {}

// DEPENDENCIES
type Handler struct {
	store Store
}

// CONSTRUCTOR
func newHandler(store Store) *Handler {
	return &Handler{store}
}
