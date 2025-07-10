package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/handlers/accounts"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/handlers/transactions"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/handlers/users"
	customMiddleware "github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/middleware"
)

func SetupRoutes(database *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	userHandler := users.UserHandler{DB: database}
	accountHandler := accounts.AccountHandler{DB: database}
	transactionHandler := transactions.TransactionsHandler{DB: database}

	r.Route("/api", func(r chi.Router) {
		r.Use(customMiddleware.RequestLogger)

		// Users
		r.Get("/users", userHandler.GetAllUsers)
		r.Get("/users/{id}", userHandler.GetUserById)
		r.Post("/users", userHandler.CreateUser)
		r.Put("/users/{id}", userHandler.UpdateUserById)
		r.Delete("/users/{id}", userHandler.DeleteUser)

		// Accounts
		r.Post("/accounts", accountHandler.CreateAccount)
		r.Get("/accounts/{user_id}", accountHandler.GetAllAccountByUserId)
		r.Delete("/accounts/{id}", accountHandler.DeleteAccount)

		// Transactions
		r.Post("/transactions", transactionHandler.CreateTransaction)
		r.Get("/transactions/{from_account_id}", transactionHandler.GetByFromAccountId)
	})

	return r
}
