package models

import "errors"

type Account struct {
	ID          int     `db:"id" json:"id"`
	UserID      *int    `db:"user_id" json:"user_id"`
	AccountName string  `db:"account_name" json:"account_name"`
	Balance     float64 `db:"balance" json:"balance"`
}

func (a *Account) Validate() error {
	if a.AccountName == "" {
		return errors.New("account_name is required")
	}

	if a.Balance == 0 {
		return errors.New("balance must be non-zero")
	}

	if a.UserID == nil {
		return errors.New("user_id is required")
	}

	return nil
}
