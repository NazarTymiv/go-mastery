package transactions

import "github.com/jmoiron/sqlx"

type TransactionHandler struct {
	DB *sqlx.DB
}
