package middleware

import (
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/db"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/helpers"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/logger"
)

func AuthError(w http.ResponseWriter) {
	logger.Error("[Authorize middleware]: Unauthorized", nil)
	helpers.SendError(w, "Unauthorized", http.StatusUnauthorized)
}

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		user, ok := db.Users[username]
		if !ok {
			AuthError(w)
			return
		}

		st, err := r.Cookie("session_token")
		if err != nil || st.Value == "" || st.Value != user.SessionToken {
			AuthError(w)
			return
		}

		csrf := r.Header.Get("X-CSRF-Token")
		if csrf != user.CSRFToken || csrf == "" {
			AuthError(w)
			return
		}

		next(w, r)
	}
}
