// Package basic contains basic get and post endpoints
package basic

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi"
)

// User entity
type User struct {
	UserID   string `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Routes returns all routes for basic api
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{userID}", GetUser)
	router.Delete("/{userID}", DeleteUser)
	router.Post("/", CreateUser)
	router.Get("/", GetAllUsers)
	return router
}

// GetUser returns a simple message to user
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	users := User{
		UserID:   userID,
		Username: "GoUser",
		Email:    "gouser@gouser.com",
	}

	render.JSON(w, r, users)
}

// CreateUser adds user to persistence
func CreateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	resp := make(map[string]string)
	msg := fmt.Sprintf("User %s created successfully", userID)
	resp["message"] = msg

	render.JSON(w, r, resp)
}

// DeleteUser deletes user from persistence
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	resp := make(map[string]string)
	msg := fmt.Sprintf("User %s deleted successfully", userID)
	resp["message"] = msg

	render.JSON(w, r, resp)
}

// GetAllUsers gets all users in persistence
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			UserID:   "userID1",
			Username: "GoUser",
			Email:    "gouser@gouser.com",
		},
	}

	render.JSON(w, r, users)
}
