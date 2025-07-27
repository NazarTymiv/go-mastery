package users

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.SendError(w, "Invalid user ID", http.StatusBadRequest)
		logger.Error("[Delete User Handler]: Invalid user ID", err.Error())
		return
	}

	statusCode, err := models.DeleteUserByID(h.DB, id)
	if err != nil {
		helpers.SendError(w, err.Error(), statusCode)
		logger.Error("[Delete User Handler]: Could not delete user by id", err.Error())
		return
	}

	helpers.SendSuccess(w, helpers.MessageResponse{Message: "User has been deleted successfully"}, http.StatusOK)
}
