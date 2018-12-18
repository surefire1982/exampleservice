// Package basic contains basic get and post endpoints
package basic

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Routes returns all routes for basic api
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetWelcome)

	return router
}

// GetWelcome returns a simple message to user
func GetWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to Go!"))
}
