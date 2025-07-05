package accounts

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
)

func (h AccountHandler) GetAllAccountByUserId(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")

	var accounts []models.Account
	err := h.DB.Select(&accounts, "SELECT * FROM accounts WHERE user_id = ?", userId)
	if err != nil {
		log.Fatalf("Error fetch user accounts: %v", err)
		helpers.WriteError(w, "Could not get user's accounts", http.StatusInternalServerError)
		return
	}

	if accounts == nil {
		helpers.WriteError(w, "Could not find any accounts related to this user", http.StatusNotFound)
		return
	}

	helpers.SendJson(w, accounts)
}
