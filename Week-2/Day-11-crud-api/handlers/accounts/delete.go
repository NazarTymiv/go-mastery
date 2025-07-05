package accounts

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/helpers"
)

func (h AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.DB.Exec("DELETE FROM accounts WHERE id = ?", id)
	if err != nil {
		helpers.WriteError(w, "Could not delete account", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("RowsAffected error: %v", err)
		helpers.WriteError(w, "Could not verify delete account", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		helpers.WriteError(w, "Could not find user with provided ID", http.StatusNotFound)
		return
	}

	helpers.SendJson(w, map[string]string{"message": "Account has bee successfully deleted"})
}
