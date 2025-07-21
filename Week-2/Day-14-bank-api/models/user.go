package models

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// SQL requests
const (
	GetUserByEmailSQL = `SELECT * FROM users WHERE email = ?`
	CreateUserSQL     = `INSERT INTO users (name, email) VALUES(:name, :email)`
	DeleteUserByIdSQL = `DELETE FROM users WHERE id = ?`
	GetAllUsersSql    = `SELECT * FROM users`
)

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func GetAllUsers(users *[]User, db *sqlx.DB) error {
	err := db.Select(users, GetAllUsersSql)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func GetUserByEmail(db *sqlx.DB, email string) (*User, error) {
	var user User
	err := db.Get(&user, GetUserByEmailSQL, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func DeleteUserByID(db *sqlx.DB, id int) (int, error) {
	res, err := db.Exec(DeleteUserByIdSQL, id)
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	if rowsAffected == 0 {
		return http.StatusNotFound, errors.New("couldn't find user with provided id")
	}

	return 0, nil
}
