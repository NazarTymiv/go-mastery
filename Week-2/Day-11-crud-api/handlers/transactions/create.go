package transactions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
)

func (h TransactionsHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.WriteError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newTransaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&newTransaction); err != nil {
		helpers.WriteError(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := newTransaction.Validate(); err != nil {
		helpers.WriteError(w, "Missing fields", http.StatusBadRequest)
		return
	}

	tx, err := h.DB.Beginx()
	if err != nil {
		helpers.WriteError(w, "Could not start transaction", http.StatusInternalServerError)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var fromAccount models.Account
	err = tx.Get(&fromAccount, "SELECT * FROM accounts WHERE id = ?", newTransaction.FromAccountId)
	if err != nil {
		helpers.WriteError(w, "Could not find provided 'From account'", http.StatusNotFound)
		return
	}

	if float64(fromAccount.Balance) < float64(newTransaction.Amount) {
		helpers.WriteError(w, fmt.Sprintf("%s does not have enough balance to send", fromAccount.AccountName), http.StatusBadRequest)
		return
	}

	var toAccount models.Account
	err = tx.Get(&toAccount, "SELECT * FROM accounts WHERE id = ?", newTransaction.ToAccountId)
	if err != nil {
		helpers.WriteError(w, "Could not find provided 'To account'", http.StatusNotFound)
		return
	}

	fromAccount.Balance -= newTransaction.Amount
	toAccount.Balance += newTransaction.Amount

	_, err = tx.NamedExec("UPDATE accounts SET balance = :balance WHERE id = :id", &fromAccount)
	if err != nil {
		log.Printf("Error updating from_account: %v", err)
		helpers.WriteError(w, fmt.Sprintf("Failed to update balance for %s", fromAccount.AccountName), http.StatusInternalServerError)
		return
	}

	_, err = tx.NamedExec("UPDATE accounts SET balance = :balance WHERE id = :id", &toAccount)
	if err != nil {
		log.Printf("Error updating to_account: %v", err)
		helpers.WriteError(w, fmt.Sprintf("Failed to update balance for %s", toAccount.AccountName), http.StatusInternalServerError)
		return
	}

	_, err = tx.NamedExec(`
        INSERT INTO transactions (from_account_id, to_account_id, amount, type, description)
        VALUES (:from_account_id, :to_account_id, :amount, :type, :description)
    `, &newTransaction)
	if err != nil {
		log.Printf("Error inserting transaction: %v", err)
		helpers.WriteError(w, "Could not create transaction", http.StatusInternalServerError)
		return
	}

	helpers.SendJson(w, newTransaction)
}
