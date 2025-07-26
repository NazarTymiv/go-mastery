package transactions

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h TransactionHandler) GetTransactionsByAccount(w http.ResponseWriter, r *http.Request) {
	accountIdStr := chi.URLParam(r, "account_id")
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		helpers.SendError(w, "Invalid account id", http.StatusBadRequest)
		logger.Error("[Get Transactions By Account Handler]: Invalid account id", err.Error())
		return
	}

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
	err = models.GetTransactionsByAccount(h.DB, &transactions, accountId, limit, offset)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Get Transactions By Account Handler]: Could not get transactions by account", err.Error())
		return
	}

	helpers.SendSuccess(w, transactions, http.StatusOK)
	logger.Info("[Get Transactions By Account Handler]: Selected all transactions by account successfully", transactions)
}
