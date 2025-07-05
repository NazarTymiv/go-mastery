package accounts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
)

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.WriteError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newAccount models.Account
	err := json.NewDecoder(r.Body).Decode(&newAccount)
	if err != nil {
		helpers.WriteError(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = newAccount.Validate()
	if err != nil {
		helpers.WriteError(w, "Missing fields", http.StatusBadRequest)
		return
	}

	var foundUser models.User
	err = h.DB.Get(&foundUser, "SELECT * FROM users WHERE id = ?", newAccount.UserID)
	if err != nil {
		helpers.WriteError(w, "Could not find user with provided ID", http.StatusNotFound)
		return
	}

	res, err := h.DB.NamedExec("INSERT INTO accounts (user_id, account_name, balance) VALUES (:user_id, :account_name, :balance)", &newAccount)
	if err != nil {
		log.Fatalf("Error insert account: %v", err)
		helpers.WriteError(w, "Could not create account", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	newAccount.ID = int(id)

	helpers.SendJson(w, newAccount)
}
