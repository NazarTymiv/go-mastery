package accounts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendError(w, "Invalid account id", http.StatusBadRequest)
		logger.Error("[Delete Account handler]: invalid account id", err.Error())
		return
	}

	statusCode, err := models.DeleteAccountById(h.DB, id)
	if err != nil {
		helpers.SendError(w, err.Error(), statusCode)
		return
	}

	helpers.SendSuccess(w, helpers.MessageResponse{Message: fmt.Sprintf("Account with id %v has been deleted successfully", id)}, statusCode)

}
