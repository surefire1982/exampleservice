// Package basic contains basic get and post endpoints
package basic

import (
	"fmt"
	"net/http"

	"github.com/surefire1982/exampleservice/pkg/entity"

	"github.com/go-chi/render"
	"github.com/surefire1982/exampleservice/internal/config"

	"github.com/go-chi/chi"
)

// Routes returns all routes for basic api
func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{userID}", GetUser)
	router.Delete("/{userID}", DeleteUser)
	router.Post("/", CreateUser)
	router.Get("/", GetAllUsers)
	return router
}

// GetUser returns a simple message to user.  GetUser Handler is a closure which accepts the configuration
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	users := entity.User{
		UserID:   userID,
		Username: "GoUser",
		Email:    "gouser@gouser.com",
	}

	render.JSON(w, r, users)
}

// CreateUser adds user to persistence.  CreateUser Handler is a closure that accepts a configuration
func CreateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	resp := make(map[string]string)
	msg := fmt.Sprintf("User %s created successfully", userID)
	resp["message"] = msg

	render.JSON(w, r, resp)
}

// DeleteUser deletes user from persistence. DeleteUser Handler is a closer that accepts the configuration
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	resp := make(map[string]string)
	msg := fmt.Sprintf("User %s deleted successfully", userID)
	resp["message"] = msg

	render.JSON(w, r, resp)
}

// GetAllUsers gets all users in persistence.  GetAllUsers Handler is a closer that accepts configuration
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []entity.User{
		{
			UserID:   "userID1",
			Username: "GoUser",
			Email:    "gouser@gouser.com",
		},
	}

	render.JSON(w, r, users)
}
