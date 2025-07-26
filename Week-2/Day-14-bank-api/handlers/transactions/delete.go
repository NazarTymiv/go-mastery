package transactions

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendError(w, "Invalid transaction id", http.StatusBadRequest)
		logger.Error("[Delete Transaction Handler]: Invalid transaction id", err.Error())
		return
	}

	statusCode, err := models.DeleteTransaction(h.DB, id)
	if err != nil {
		helpers.SendError(w, err.Error(), statusCode)
		return
	}

	helpers.SendSuccess(w, helpers.MessageResponse{Message: fmt.Sprintf("Transaction with id %v has been deleted successfully", id)}, statusCode)
	logger.Info("[Delete Transaction Handler]: Transaction has been deleted successfully", id)
}
