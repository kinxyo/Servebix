package user

import (
	"github.com/kinxyo/Servebix.git/dashboard"
	"github.com/kinxyo/Servebix.git/utils"
	"log"
	"net/http"
)

type Routes interface {
	handleLogIn(w http.ResponseWriter, r *http.Request)
	handleSignUp(w http.ResponseWriter, r *http.Request)
}

// ACTIVATE METHODS
func (handler *Handler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("POST /patient/signup", handler.handleSignUp)
	router.HandleFunc("POST /patient/login", handler.handleLogIn)
	router.HandleFunc("GET /dashboard/users", handler.userDashboard)
}

// METHODS
func (handler *Handler) userDashboard(writer http.ResponseWriter, _ *http.Request) {
	dashboard.ShowAll(writer, handler.store.db)
}

func (handler *Handler) handleSignUp(w http.ResponseWriter, r *http.Request) {

	log.Println("Initiating sign up")

	// Parse payload
	var user Credentials
	err := utils.ParsePayload(r.Body, &user)
	if err != nil {
		log.Printf("Error parsing payload: %v", err)
		utils.RespondError(w, 500, err)
		return
	}

	// Create User
	err = handler.store.RegisterUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		utils.RespondError(w, 500, err)
		return
	}

	log.Println("User created!")

	// Return response
	utils.Respond(w, http.StatusCreated, nil)
}

func (handler *Handler) handleLogIn(w http.ResponseWriter, r *http.Request) {

	// Parse payload
	var user Credentials
	err := utils.ParsePayload(r.Body, &user)
	if err != nil {
		log.Printf("Error parsing payload: %v", err)
		utils.RespondError(w, 500, err)
		return
	}

	// Validate User
	var session Session
	session, err = handler.store.ValidateUser(user)
	if err != nil {
		log.Printf("Error validating user: %v", err)
		utils.RespondError(w, 500, err)
		return
	}

	// Return Session
	utils.Respond(w, http.StatusOK, session)
}

// DEPENDENCIES
type Handler struct {
	store *Store
}

// CONSTRUCTOR
func NewHandler(store *Store) *Handler {
	return &Handler{store}
}
