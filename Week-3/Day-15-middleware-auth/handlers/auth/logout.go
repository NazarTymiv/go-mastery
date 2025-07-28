package auth

import (
	"net/http"
	"time"

	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/db"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/helpers"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	username := r.FormValue("username")
	user, _ := db.Users[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	db.Users[username] = user

	helpers.SendSuccess(w, helpers.MessageResponse{Message: "Logged out successfully!"}, http.StatusOK)
}
