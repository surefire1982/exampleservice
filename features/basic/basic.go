// Package basic contains basic get and post endpoints
package basic

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/surefire1982/exampleservice/internal/config"

	"github.com/go-chi/chi"
)

// Config entity
type Config struct {
	*config.Config
}

// New configuration
func New(configuration *config.Config) *Config {
	return &Config{configuration}
}

// User entity
type User struct {
	UserID   string `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Routes returns all routes for basic api
func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{userID}", GetUser(configuration))
	router.Delete("/{userID}", DeleteUser(configuration))
	router.Post("/", CreateUser(configuration))
	router.Get("/", GetAllUsers(configuration))
	return router
}

// GetUser returns a simple message to user.  GetUser Handler is a closure which accepts the configuration
func GetUser(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		users := User{
			UserID:   userID,
			Username: "GoUser",
			Email:    "gouser@gouser.com",
		}

		render.JSON(w, r, users)
	}
}

// CreateUser adds user to persistence.  CreateUser Handler is a closure that accepts a configuration
func CreateUser(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		resp := make(map[string]string)
		msg := fmt.Sprintf("User %s created successfully", userID)
		resp["message"] = msg

		render.JSON(w, r, resp)
	}
}

// DeleteUser deletes user from persistence. DeleteUser Handler is a closer that accepts the configuration
func DeleteUser(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		resp := make(map[string]string)
		msg := fmt.Sprintf("User %s deleted successfully", userID)
		resp["message"] = msg

		render.JSON(w, r, resp)
	}
}

// GetAllUsers gets all users in persistence.  GetAllUsers Handler is a closer that accepts configuration
func GetAllUsers(configuration *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{
				UserID:   "userID1",
				Username: "GoUser",
				Email:    "gouser@gouser.com",
			},
		}

		render.JSON(w, r, users)
	}
}
