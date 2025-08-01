package transactions

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
)

func (h TransactionsHandler) GetByFromAccountId(w http.ResponseWriter, r *http.Request) {
	fromAccountId := chi.URLParam(r, "from_account_id")

	var transactions []models.Transaction
	err := h.DB.Select(&transactions, "SELECT * FROM transactions WHERE from_account_id = ?", fromAccountId)
	if err != nil {
		log.Fatalf("%v", err)
		helpers.WriteError(w, "Could not find transaction for provided account", http.StatusNotFound)
		return
	}

	if transactions == nil {
		helpers.WriteError(w, "Could not find any transaction for this account", http.StatusNotFound)
		return
	}

	helpers.SendJson(w, transactions)
}
