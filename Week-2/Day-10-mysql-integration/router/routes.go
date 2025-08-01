package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/handlers/users"
)

func SetupRoutes(database *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userHandler := users.UserHandler{DB: database}

	r.Route("/api", func(r chi.Router) {
		r.Get("/users", userHandler.GetUsers)
		r.Get("/users/{id}", userHandler.GetUserById)
		r.Post("/users", userHandler.CreateUser)
		r.Put("/users/{id}", userHandler.UpdateUser)
		r.Delete("/users/{id}", userHandler.DeleteUser)
	})

	return r
}
