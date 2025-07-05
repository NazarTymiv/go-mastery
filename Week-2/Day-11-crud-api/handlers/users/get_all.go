package users

import (
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
)

func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	err := h.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Fatalf("Error selecting users: %v", err)
		helpers.WriteError(w, "Could not fetch users", http.StatusInternalServerError)
		return
	}

	helpers.SendJson(w, users)
}
