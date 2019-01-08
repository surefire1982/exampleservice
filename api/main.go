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

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Routes generates routes for service
func Routes(configuration *config.Config, db *gorm.DB) *chi.Mux {
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
	userRepo := pkguser.NewDBRepository(db)
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

	dbArgs := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.Constants.DBUser, configuration.Constants.DBPassword, configuration.Constants.DBName)
	db, err := gorm.Open("mysql", dbArgs)
	defer db.Close() // defer this operation to when we kill service
	router := Routes(configuration, db)
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
