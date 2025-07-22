package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/models"
)

func (h UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		helpers.SendError(w, "Invalid user id", http.StatusBadRequest)
		logger.Error("[Update User Handler]: Invalid user id", err.Error())
		return
	}

	// Getting Update User data
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		helpers.SendError(w, "Invalid body", http.StatusBadRequest)
		logger.Error("[Update User Handler]: Invalid request body", err.Error())
		return
	}

	// Validating if all the fields in the body exists
	err = updatedUser.Validate()
	if err != nil {
		helpers.SendError(w, err.Error(), http.StatusBadRequest)
		logger.Error("[Update User Handler]: missing fields (both name and email are required)", err.Error())
		return
	}

	// Checking if email already in use
	foundUser, err := models.GetUserByEmail(h.DB, updatedUser.Email)
	if err != nil {
		helpers.SendError(w, "Server Error", http.StatusInternalServerError)
		logger.Error("[Update User Handler]: Could not get user by email from db", err.Error())
		return
	}

	if foundUser != nil && foundUser.ID != userId {
		helpers.SendError(w, "Email already in use", http.StatusConflict)
		logger.Error("[Update User Handler]: Provided email of user already exists", foundUser.Email)
		return
	}

	// Updating User in DB
	updatedUser.ID = userId
	statusCode, err := models.UpdateUser(h.DB, updatedUser)
	if err != nil {
		helpers.SendError(w, err.Error(), statusCode)
		return
	}

	helpers.SendSuccess(w, helpers.MessageResponse{Message: fmt.Sprintf("User with id %v has been successfully updated", updatedUser.ID)}, statusCode)
	logger.Info("[Update User Handler]: User updated successfully", updatedUser)
}
