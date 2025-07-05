package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
)

func (h UserHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
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

	if updatedUser.Name == "" {
		helpers.WriteError(w, "Missing name", http.StatusBadRequest)
		return
	}

	updatedUser.ID = int(id)

	res, err := h.DB.NamedExec("UPDATE users SET name = :name WHERE id = :id", &updatedUser)
	if err != nil {
		log.Fatalf("Update error: %v", err)
		helpers.WriteError(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("RowsAffected error: %v", err)
		helpers.WriteError(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		helpers.WriteError(w, "User with provided id not found", http.StatusNotFound)
		return
	}

	helpers.SendJson(w, map[string]string{"message": "user updated successfully"})
}
