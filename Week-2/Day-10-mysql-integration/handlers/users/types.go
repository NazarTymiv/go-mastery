package users

import "github.com/jmoiron/sqlx"

type UserHandler struct {
	DB *sqlx.DB
}
