package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
)

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helpers.WriteError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	res, err := h.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		helpers.WriteError(w, "Could not delete user", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("RowsAffected error: %v", err)
		helpers.WriteError(w, "Could not verify user deletion", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		helpers.WriteError(w, "Could not find user with provided id", http.StatusNotFound)
		return
	}

	helpers.SendJson(w, map[string]string{"message": "User has been deleted successfully"})
}
