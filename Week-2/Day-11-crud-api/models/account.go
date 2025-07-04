package models

type Account struct {
	ID          int    `db:"id" json:"id"`
	UserID      int    `user_id:"id" json:"user_id"`
	AccountName string `db:"account_name" json:"account_name"`
	Balance     int    `db:"balance" json:"balance"`
}
