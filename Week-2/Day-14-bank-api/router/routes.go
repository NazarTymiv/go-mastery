package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/handlers/accounts"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/handlers/status"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/handlers/transactions"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/handlers/users"
	customMiddleware "github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/middleware"
)

func SetupRoutes(database *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	usersHandler := users.UserHandler{DB: database}
	accountsHandler := accounts.AccountHandler{DB: database}
	transactionsHandler := transactions.TransactionHandler{DB: database}

	// Server Health Check
	r.Get("/ping", status.ServerHealthCheck)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use(customMiddleware.RequestLogger)

			// Users routes
			r.Get("/users", usersHandler.GetAllUsers)
			r.Post("/users", usersHandler.CreateUser)
			r.Put("/users/{id}", usersHandler.UpdateUserByID)
			r.Delete("/users/{id}", usersHandler.DeleteUser)

			// Accounts routes
			r.Get("/accounts", accountsHandler.GetAllAccounts)
			r.Get("/accounts/{user_id}", accountsHandler.GetAllByUserId)
			r.Get("/accounts/balance/{id}", accountsHandler.GetBalanceOfAccount)
			r.Post("/accounts", accountsHandler.CreateAccount)
			r.Delete("/accounts/{id}", accountsHandler.DeleteAccount)

			// Transactions routes
			r.Get("/transactions", transactionsHandler.GetAllTransactions)
			r.Get("/transactions/account/{account_id}", transactionsHandler.GetTransactionsByAccount)
			r.Post("/transactions", transactionsHandler.CreateNewTransaction)
			r.Delete("/transactions/{id}", transactionsHandler.DeleteTransaction)
		})
	})

	return r
}
