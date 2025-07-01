package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/models"
	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/response"
)

func StatusHandler(serverStart time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		countedUptime := fmt.Sprintf("%.0fs", time.Since(serverStart).Seconds())

		response.SendJSON(w, http.StatusOK, models.StatusResponse{Uptime: countedUptime, Status: "ok"})
	}
}
