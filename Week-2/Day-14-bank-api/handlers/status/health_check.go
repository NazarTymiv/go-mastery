package status

import (
	"net/http"
	"time"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/config"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/helpers"
)

type HealthCheck struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

func ServerHealthCheck(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(config.StartTime).Truncate(time.Second)

	response := HealthCheck{
		Status: "ok",
		Uptime: uptime.String(),
	}

	helpers.SendSuccess(w, response, http.StatusOK)
}
