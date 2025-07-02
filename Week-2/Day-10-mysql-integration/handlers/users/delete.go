package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/helpers"
)

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.WriteError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	result, err := h.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		log.Printf("Delete error: %v", err)
		helpers.WriteError(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("RowsAffected error: %v", err)
		helpers.WriteError(w, "Failed to confirm deletion", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		helpers.WriteError(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User has been successfully deleted"})
}
