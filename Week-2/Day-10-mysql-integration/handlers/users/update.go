package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/models"
)

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.WriteError(w, "Invalid user ID", http.StatusBadRequest)
	}

	if r.Method != http.MethodPut {
		helpers.WriteError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		helpers.WriteError(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if updatedUser.Email == "" && updatedUser.Name == "" {
		helpers.WriteError(w, "Missing field", http.StatusBadRequest)
		return
	}

	updatedUser.ID = uint8(id)

	result, err := h.DB.NamedExec("UPDATE users SET name = :name, email = :email WHERE id = :id", &updatedUser)
	if err != nil {
		log.Printf("Update error: %v", err)
		helpers.WriteError(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		helpers.WriteError(w, "No user found with given ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User has been successfully updated"})
}
