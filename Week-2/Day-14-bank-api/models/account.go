package models

type Account struct {
	ID      int     `db:"id" json:"id"`
	UserId  *int    `db:"user_id" json:"user_id"`
	Name    string  `db:"account_name" json:"account_name"`
	Balance float64 `db:"balance" json:"balance"`
}
