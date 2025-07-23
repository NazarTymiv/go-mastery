package users

import (
	"net/http"
	"strconv"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 10

	queryPage := r.URL.Query().Get("page")
	queryLimit := r.URL.Query().Get("limit")

	if queryLimit != "" {
		if p, err := strconv.Atoi(queryPage); err == nil && p > 0 {
			page = p
		}
	}

	if queryPage != "" {
		if l, err := strconv.Atoi(queryLimit); err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	users := []models.User{}
	err := models.GetAllUsers(&users, h.DB, limit, offset)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Get All Users Handler]: Could not get any users", err.Error())
		return
	}

	helpers.SendSuccess(w, users, http.StatusOK)
}
