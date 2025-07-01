package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/models"
)

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user models.User
	err := h.DB.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Printf("Query error: %v", err)
		http.Error(w, "User with provided id not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
