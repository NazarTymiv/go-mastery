package transactions

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
)

func (h TransactionsHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.WriteError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTransaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&newTransaction)
	if err != nil {
		helpers.WriteError(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = newTransaction.Validate()
	if err != nil {
		helpers.WriteError(w, "Missing fields", http.StatusBadRequest)
		return
	}

	var foundAccount models.Account
	err = h.DB.Get(&foundAccount, "SELECT * FROM accounts WHERE id = ?", newTransaction.AccountID)
	if err != nil {
		helpers.WriteError(w, "Could not find provided account", http.StatusNotFound)
		return
	}

	newTransaction.CreatedAt = time.Now()

	res, err := h.DB.NamedExec("INSERT INTO transactions (account_id, amount, type, description, created_at) VALUES (:account_id, amount, type, description, created_at)", &newTransaction)
	if err != nil {
		log.Fatalf("Error create transaction: %v", err)
		helpers.WriteError(w, "Could not create transaction", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	newTransaction.ID = int(id)

	helpers.SendJson(w, newTransaction)
}
