package auth

import (
	"fmt"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/db"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/helpers"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/logger"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/models"
	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) < 8 || len(password) < 8 {
		helpers.SendError(w, "Invalid username/password", http.StatusNotAcceptable)
		logger.Error("[Register Handler]: Invalid username/password", fmt.Sprintf("username: %s, password: %s", username, password))
		return
	}

	if _, ok := db.Users[username]; ok {
		helpers.SendError(w, "User already exists", http.StatusConflict)
		logger.Error("[Register Handler]: User already exists", username)
		return
	}

	hashedPassword, _ := utils.HashPassword(password)
	db.Users[username] = models.Login{
		HashedPassword: hashedPassword,
	}

	helpers.SendSuccess(w, db.Users, http.StatusCreated)
}
