package models

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func GetUserByEmail(db *sqlx.DB, email string) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
