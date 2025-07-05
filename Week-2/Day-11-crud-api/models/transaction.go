package models

import (
	"errors"
	"time"
)

type Transaction struct {
	ID          int       `db:"id" json:"id"`
	AccountID   *int      `db:"account_id" json:"account_id"`
	Amount      float64   `db:"amount" json:"amount"`
	Type        string    `db:"type" json:"type"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

func (t *Transaction) Validate() error {
	if t.AccountID == nil {
		return errors.New("account_id required")
	}

	if t.Amount == 0 {
		return errors.New("amount can not to be zero")
	}

	if t.Type == "" {
		return errors.New("type required")
	}

	if t.Description == "" {
		return errors.New("description required")
	}

	return nil
}
