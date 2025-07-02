package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/models"
)

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.WriteError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		helpers.WriteError(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if newUser.Email == "" || newUser.Name == "" {
		helpers.WriteError(w, "Missing fields", http.StatusBadRequest)
		return
	}

	result, err := h.DB.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", &newUser)
	if err != nil {
		log.Printf("Insert error: %v", err)
		helpers.WriteError(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	newUser.ID = uint8(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}
