package accounts

import (
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
)

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	helpers.SendSuccess(w, helpers.MessageResponse{Message: "Create Account route"}, http.StatusCreated)
}
