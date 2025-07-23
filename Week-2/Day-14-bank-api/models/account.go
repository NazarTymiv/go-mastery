package models

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
)

type Account struct {
	ID      int      `db:"id" json:"id"`
	UserId  *int     `db:"user_id" json:"user_id"`
	Name    string   `db:"account_name" json:"account_name"`
	Balance *float64 `db:"balance" json:"balance"`
}

// SQL requests
const (
	CreateNewAccountSQL = `INSERT INTO accounts (user_id, account_name, balance) VALUES(:user_id, :account_name, :balance)`
)

func (a *Account) Validate() error {
	if a.UserId == nil {
		return errors.New("user id is required")
	}

	if a.Name == "" {
		return errors.New("name is required")
	}

	if a.Balance == nil {
		return errors.New("balance is required")
	}

	return nil
}

func CreateNewAccount(db *sqlx.DB, newAccount *Account) error {
	res, err := db.NamedExec(CreateNewAccountSQL, &newAccount)
	if err != nil {
		logger.Error("[Create New Account DB]: Could not create new account", err.Error())
		return errors.New("server error")
	}

	accountId, err := res.LastInsertId()
	if err != nil {
		logger.Error("[Create New Account DB]: Could not verify created account", err.Error())
		return errors.New("server error")
	}

	newAccount.ID = int(accountId)
	return nil
}
