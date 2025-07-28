package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/handlers/auth"
	customMiddleware "github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/middleware"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Use(customMiddleware.RequestLogger)

		r.Post("/register", auth.Register)
		r.Post("/login", auth.Login)
		r.Post("/logout", customMiddleware.Authorize(auth.Logout))
		r.Get("/protected", customMiddleware.Authorize(auth.Protected))
	})

	return r
}
