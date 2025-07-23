package accounts

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h AccountHandler) GetBalanceOfAccount(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendError(w, "Invalid account id", http.StatusBadRequest)
		logger.Error("[Delete Account handler]: invalid account id", err.Error())
		return
	}

	// Find account by id
	foundAccount, err := models.GetAccountById(h.DB, id)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Get Balance of Account]: Could not get balance of account", err.Error())
		return
	}

	if foundAccount == nil {
		helpers.SendError(w, "Could not find account with provided id", http.StatusNotFound)
		logger.Error("[Get Balance of Account]: Could not found account with provided id", nil)
		return
	}

	helpers.SendSuccess(w, map[string]float64{"balance": *foundAccount.Balance}, http.StatusOK)
	logger.Info("[Get Balance of Account]: Successfully got balance of account", foundAccount.Balance)
}
