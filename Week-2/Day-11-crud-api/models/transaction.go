package models

type Transaction struct {
	ID          int     `db:"id" json:"id"`
	AccountID   int     `db:"account_id" json:"account_id"`
	Amount      float64 `db:"amount" json:"amount"`
	Type        string  `db:"type" json:"type"`
	Description string  `db:"description" json:"description"`
	CreatedAt   string  `db:"created_at" json:"created_at"`
}
