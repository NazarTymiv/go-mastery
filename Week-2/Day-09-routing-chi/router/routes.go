package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/handlers"
	customMiddleware "github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/middleware"
)

func SetupRoutes(serverStart time.Time) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(customMiddleware.RequestTimer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", handlers.PingHandler)
		r.Get("/greet", handlers.GreetHandler)
		r.Post("/echo", handlers.EchoHandler)
		r.Get("/status", handlers.StatusHandler(serverStart))
	})

	return r
}
