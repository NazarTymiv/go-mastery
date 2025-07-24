package models

import (
	"errors"
	"time"
)

type TransactionType string

const (
	Deposit    TransactionType = "deposit"
	Withdrawal TransactionType = "withdrawal"
	Transfer   TransactionType = "transfer"
)

type Transaction struct {
	ID          int             `db:"id" json:"id"`
	Sender      *int            `db:"from_account_id" json:"from_account_id"`
	Receiver    *int            `db:"to_account_id" json:"to_account_id"`
	Amount      *float64        `db:"amount" json:"amount"`
	Type        TransactionType `db:"type" json:"type"`
	Description string          `db:"description" json:"description"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
}

func (t *Transaction) Validate() error {
	if t.Sender == nil {
		return errors.New("sender id is required")
	}

	if t.Receiver == nil {
		return errors.New("receiver id is required")
	}

	if t.Amount == nil {
		return errors.New("amount is required")
	}

	if t.Type != Deposit && t.Type != Withdrawal && t.Type != Transfer {
		return errors.New("transaction type is required and has to be: deposit, withdrawal or transfer")
	}

	if t.Description == "" {
		return errors.New("description is required")
	}

	return nil
}
