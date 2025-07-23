package models

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
)

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// SQL requests
const (
	GetAllUsersSql    = `SELECT * FROM users ORDER BY id LIMIT ? OFFSET ?`
	GetUserByEmailSQL = `SELECT * FROM users WHERE email = ?`
	CreateUserSQL     = `INSERT INTO users (name, email) VALUES(:name, :email)`
	UpdateUserSQL     = `UPDATE users SET name = :name, email = :email WHERE id = :id`
	DeleteUserByIdSQL = `DELETE FROM users WHERE id = ?`
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

func GetAllUsers(users *[]User, db *sqlx.DB, limit int, offset int) error {
	err := db.Select(users, GetAllUsersSql, limit, offset)
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

func UpdateUser(db *sqlx.DB, updatedUser User) (int, error) {
	updatedUser.Email = strings.TrimSpace(strings.ToLower(updatedUser.Email))
	res, err := db.NamedExec(UpdateUserSQL, updatedUser)
	if err != nil {
		logger.Error("[Update User]: Could not update user", err.Error())
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Error("[Update User]: Could not verify updated user", err.Error())
		return http.StatusInternalServerError, errors.New("internal server error")
	}

	if rowsAffected == 0 {
		logger.Error("[Update User]: Could not find user with given ID", nil)
		return http.StatusNotFound, errors.New("no user found with given ID")
	}

	return http.StatusOK, nil
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
