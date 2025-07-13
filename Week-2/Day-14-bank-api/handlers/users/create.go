package users

import (
	"encoding/json"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.SendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		logger.Error("[Create User Handler]: Method not allowed", nil)
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		helpers.SendError(w, "Invalid body", http.StatusBadRequest)
		logger.Error("[Create User Handler]: Invalid body", err.Error())
		return
	}

	err = newUser.Validate()
	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusBadRequest)
		logger.Error("[Create User Handler]: Missing fields", err.Error())
		return
	}

	foundUser, err := models.GetUserByEmail(h.DB, newUser.Email)
	if err != nil {
		helpers.SendError(w, "Could not check if user exist", http.StatusInternalServerError)
		logger.Error("[Create User Handler]: Could not get user by email from db", err.Error())
		return
	}

	if foundUser != nil {
		helpers.SendError(w, "Email already in use", http.StatusBadRequest)
		logger.Error("[Create User Handler]: Provided email of user already exists", foundUser.Email)
		return
	}

	res, err := h.DB.NamedExec("INSERT INTO users (name, email) VALUES(:name, :email)", &newUser)
	if err != nil {
		helpers.SendError(w, "Could not create user", http.StatusInternalServerError)
		logger.Error("[Create User Handler]: Could not create user", err.Error())
		return
	}

	userId, err := res.LastInsertId()
	if err != nil {
		helpers.SendError(w, "Could not create user", http.StatusInternalServerError)
		logger.Error("[Create User Handler]: Could get id of created user", err.Error())
		return
	}

	newUser.ID = int(userId)

	helpers.SendSuccess(w, newUser, http.StatusCreated)
}
