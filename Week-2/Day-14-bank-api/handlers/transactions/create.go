package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h TransactionHandler) CreateNewTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&newTransaction)
	if err != nil {
		helpers.SendError(w, "Invalid body", http.StatusBadRequest)
		logger.Error("[Create New Transaction Handler]: Invalid body", err.Error())
		return
	}

	// Validate Transaction body
	err = newTransaction.Validate()
	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusBadRequest)
		logger.Error("[Create New Transaction Handler]: Missing fields", err.Error())
		return
	}

	// Getting sender and receiver accounts
	senderAccount, err := models.GetAccountById(h.DB, newTransaction.Sender)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Create New Transaction Handler]: Could not get sender account", err.Error())
		return
	}

	if senderAccount == nil {
		helpers.SendError(w, "Could not find sender account with provided id", http.StatusNotFound)
		logger.Error("[Create New Transaction Handler]: could not find sender account", nil)
		return
	}

	receiverAccount, err := models.GetAccountById(h.DB, newTransaction.Receiver)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Create New Transaction Handler]: Could not get receiver account", err.Error())
		return
	}

	if receiverAccount == nil {
		helpers.SendError(w, "Could not find receiver account with provided id", http.StatusNotFound)
		logger.Error("[Create New Transaction Handler]: could not find receiver account", nil)
		return
	}

	//
}
