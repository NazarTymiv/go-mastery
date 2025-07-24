package transactions

import (
	"encoding/json"
	"fmt"
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

	// Begin transaction
	statusCode, err := models.CreateTransaction(h.DB, &newTransaction)
	if err != nil {
		helpers.SendError(w, err.Error(), statusCode)
		return
	}

	helpers.SendSuccess(w, helpers.MessageResponse{Message: fmt.Sprintf("Account with id %v successfully sent £%.2f to account with id %v", *newTransaction.Sender, *newTransaction.Amount, *newTransaction.Receiver)}, statusCode)
}
