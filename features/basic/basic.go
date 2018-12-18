// Package basic contains basic get and post endpoints
package basic

import (
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
