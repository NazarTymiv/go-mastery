package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/models"
)

func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	err := h.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		log.Printf("Query error: %v", err)
		helpers.WriteError(w, "Could not fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
