package users

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/models"
)

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user models.User
	err := h.DB.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		log.Printf("Error selecting user by id: %v", err)
		helpers.WriteError(w, "Could not find user with provided id", http.StatusNotFound)
		return
	}

	helpers.SendJson(w, user)
}
