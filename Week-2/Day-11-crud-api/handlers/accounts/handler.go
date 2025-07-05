package accounts

import "github.com/jmoiron/sqlx"

type AccountHandler struct {
	DB *sqlx.DB
}
