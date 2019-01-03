package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	user "github.com/surefire1982/exampleservice/api/handlers/user"
	"github.com/surefire1982/exampleservice/internal/config"
	pkguser "github.com/surefire1982/exampleservice/pkg/user"
)

// Routes generates routes for service
func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	// setup handlers
	// userHandler
	userRepo := pkguser.NewInMemRepository()
	userSvc := pkguser.NewService(userRepo)
	userHandler := user.NewUserHandler(*userSvc)

	// add database connection strings here
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/user", userHandler.Routes(configuration))
	})

	return router

}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	router := Routes(configuration)
	port := fmt.Sprintf(":%s", configuration.Constants.PORT)

	log.Printf("Starting server on port %s\n", port)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // walk all routes
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err %s\n", err.Error()) // panic if there is an error
	}

	log.Fatal(http.ListenAndServe(port, router))
}
