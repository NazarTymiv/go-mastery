package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/handlers/users"
	customMiddleware "github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/middleware"
)

func SetupRoutes(database *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	usersHandler := users.UserHandler{DB: database}

	r.Route("/api", func(r chi.Router) {
		r.Use(customMiddleware.RequestLogger)

		// Users routes
		r.Get("/users", usersHandler.GetAllUsers)
		r.Post("/users", usersHandler.CreateUser)
		r.Put("/users/{id}", usersHandler.UpdateUserByID)
		r.Delete("/users/{id}", usersHandler.DeleteUser)
	})

	return r
}
