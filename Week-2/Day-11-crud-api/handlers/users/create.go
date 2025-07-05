package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
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

	if newUser.Name == "" {
		helpers.WriteError(w, "Missing name", http.StatusBadRequest)
		return
	}

	res, err := h.DB.NamedExec("INSERT INTO users (name) VALUES (:name)", &newUser)
	if err != nil {
		log.Printf("Error Inserting created user into db: %v", err)
		helpers.WriteError(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	newUser.ID = int(id)

	helpers.SendJson(w, newUser)
}
