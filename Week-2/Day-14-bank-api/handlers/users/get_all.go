package users

import (
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	err := models.GetAllUsers(&users, h.DB)
	if err != nil {
		helpers.SendError(w, "Server error", http.StatusInternalServerError)
		logger.Error("[Get All Users Handler]: Could not get any users", err.Error())
		return
	}

	helpers.SendSuccess(w, users, http.StatusOK)
}
