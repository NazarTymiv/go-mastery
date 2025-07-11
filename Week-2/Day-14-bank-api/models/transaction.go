package models

import "time"

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
	Amount      float64         `db:"amount" json:"amount"`
	Type        TransactionType `db:"type" json:"type"`
	Description string          `db:"description" json:"description"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
}
