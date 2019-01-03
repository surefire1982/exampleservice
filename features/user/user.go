// Package basic contains basic get and post endpoints
package basic

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/surefire1982/exampleservice/pkg/entity"
	"github.com/surefire1982/exampleservice/pkg/user"

	"github.com/go-chi/render"
	"github.com/surefire1982/exampleservice/internal/config"

	"github.com/go-chi/chi"
)

// UserHandler to handle user requests
type UserHandler struct {
	userSvc user.Service
}

// NewUserHandler create new handler
func NewUserHandler(userSvc user.Service) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

// Routes returns all routes for basic api
func (handler *UserHandler) Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{userID}", handler.Find)
	router.Delete("/{userID}", handler.Delete)
	router.Post("/", handler.Create)
	router.Get("/", handler.FindAll)
	return router
}

// Find a user
func (handler *UserHandler) Find(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	user, err := handler.userSvc.Find(userID)
	if err != nil {
		if err == entity.ErrNotFound {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		log.Printf("Unknown Error: %s\n", err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	render.JSON(w, r, user)
}

// Create a user
func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser entity.User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userID, err := handler.userSvc.Store(&newUser)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	log.Printf("New User Added: %s\n", userID)

	resp := entity.UserResponse{
		UserID: userID,
	}
	render.JSON(w, r, resp)
}

// Delete a user
func (handler *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	err := handler.userSvc.Delete(userID)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	resp := entity.UserResponse{
		UserID: userID,
	}
	render.JSON(w, r, resp)
}

// FindAll users
func (handler *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := handler.userSvc.FindAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(users) == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	render.JSON(w, r, users)
}
