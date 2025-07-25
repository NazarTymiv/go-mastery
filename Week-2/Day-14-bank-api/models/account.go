package models

import (
	"database/sql"
	"errors"
	"net/http"

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
	GetAllAccountsSQL         = `SELECT * FROM accounts ORDER BY id LIMIT ? OFFSET ?`
	GetAllAccountsByUserIdSQL = `SELECT * FROM accounts WHERE user_id = ? ORDER BY id`
	GetAccountByIdSQL         = `SELECT * FROM accounts WHERE id = ?`
	CreateNewAccountSQL       = `INSERT INTO accounts (user_id, account_name, balance) VALUES(:user_id, :account_name, :balance)`
	DeleteAccountByIdSQL      = `DELETE FROM accounts WHERE id = ?`
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

func (a *Account) GetBalance() float64 {
	return *a.Balance
}

func GetAllAccounts(db *sqlx.DB, accounts *[]Account, limit int, offset int) error {
	err := db.Select(accounts, GetAllAccountsSQL, limit, offset)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func GetAllAccountsByUserId(db *sqlx.DB, accounts *[]Account, userId int) error {
	err := db.Select(accounts, GetAllAccountsByUserIdSQL, userId)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func GetAccountById(db *sqlx.DB, accountId *int) (*Account, error) {
	var account Account
	err := db.Get(&account, GetAccountByIdSQL, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &account, nil
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

func DeleteAccountById(db *sqlx.DB, id int) (int, error) {
	res, err := db.Exec(DeleteAccountByIdSQL, id)
	if err != nil {
		logger.Error("[Delete Account DB]: Could not delete account", err.Error())
		return http.StatusInternalServerError, errors.New("server error")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Error("[Delete Account DB]: Could not verify deletion account", err.Error())
		return http.StatusInternalServerError, errors.New("server error")
	}

	if rowsAffected == 0 {
		logger.Error("[Delete Account DB]: Could not find account with provided id", nil)
		return http.StatusNotFound, errors.New("could not found account")
	}

	return http.StatusOK, nil
}
