package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount models.Account
	err := json.NewDecoder(r.Body).Decode(&newAccount)
	if err != nil {
		helpers.SendError(w, "Invalid body", http.StatusBadRequest)
		logger.Error("[Create Account Handler]: Invalid body", err.Error())
		return
	}

	// Validate request body
	err = newAccount.Validate()
	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusBadRequest)
		logger.Error("[Create Account Handler]: Missing fields in body", err.Error())
		return
	}

	// Check if user exist
	foundUser, err := models.GetUserById(h.DB, *newAccount.UserId)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Create Account Handler]: Failed to fetch user", err.Error())
		return
	}

	if foundUser == nil {
		helpers.SendError(w, "User not found", http.StatusNotFound)
		logger.Error("[Create Account Handler]: No user with provided user_id", newAccount.UserId)
		return
	}

	// Creating new account for user
	err = models.CreateNewAccount(h.DB, &newAccount)
	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.SendSuccess(w, newAccount, http.StatusOK)
}
