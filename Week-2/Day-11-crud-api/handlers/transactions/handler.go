package transactions

import "github.com/jmoiron/sqlx"

type TransactionsHandler struct {
	DB *sqlx.DB
}
