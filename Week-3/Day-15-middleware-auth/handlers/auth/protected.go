package auth

import (
	"fmt"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/helpers"
)

func Protected(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	helpers.SendSuccess(w, helpers.MessageResponse{Message: fmt.Sprintf("CSRF validation successful! Welcome, %s", username)}, http.StatusOK)
}
