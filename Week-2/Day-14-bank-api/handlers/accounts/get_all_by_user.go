package accounts

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h AccountHandler) GetAllByUserId(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		helpers.SendError(w, "Invalid user id", http.StatusBadRequest)
		logger.Error("[Get All By User Id Handler]: Invalid user id", err.Error())
		return
	}

	// Checking if user exists
	foundUser, err := models.GetUserById(h.DB, userId)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Get All By User Id Handler]: Failed to fetch user", err.Error())
		return
	}

	if foundUser == nil {
		helpers.SendError(w, "User not found", http.StatusNotFound)
		logger.Error("[Get All By User Id Handler]: No user with provided user_id", userId)
		return
	}

	// Selecting account by user id
	accounts := []models.Account{}
	err = models.GetAllAccountsByUserId(h.DB, &accounts, userId)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Get All By User Id Handler]: Could not select account by user id", err.Error())
		return
	}

	helpers.SendSuccess(w, accounts, http.StatusOK)
	logger.Info("[Get All By User Id Handler]: Successfully selected account by user id", accounts)
}
