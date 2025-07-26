package transactions

import (
	"net/http"
	"strconv"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h TransactionHandler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 10

	queryPage := r.URL.Query().Get("page")
	queryLimit := r.URL.Query().Get("limit")

	if queryPage != "" {
		if p, err := strconv.Atoi(queryPage); err == nil && p > 0 {
			page = p
		}
	}

	if queryLimit != "" {
		if l, err := strconv.Atoi(queryLimit); err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	transactions := []models.Transaction{}
	err := models.GetAllTransactions(h.DB, &transactions, limit, offset)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Get All Transactions Handler]: Could not get transactions", err.Error())
		return
	}

	helpers.SendSuccess(w, transactions, http.StatusOK)
	logger.Info("[Get All Transactions Handler]: Selected all transaction successfully", transactions)
}
