package auth

import (
	"net/http"
	"time"

	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/db"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/helpers"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/logger"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := db.Users[username]
	if !ok || !utils.CheckPasswordHash(password, user.HashedPassword) {
		helpers.SendError(w, "invalid username/password", http.StatusUnauthorized)
		logger.Error("[Login Handler]: Invalid credentials", map[string]string{"username": username, "password": password})
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	db.Users[username] = user

	helpers.SendSuccess(w, helpers.MessageResponse{Message: "Login successful!"}, http.StatusOK)
}
