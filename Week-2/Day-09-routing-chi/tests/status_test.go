package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/handlers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/models"
)

func TestStatusHandler(t *testing.T) {
	serverStart := time.Now().Add(-5 * time.Second)

	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	w := httptest.NewRecorder()

	handler := handlers.StatusHandler(serverStart)
	handler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

	var body models.StatusResponse
	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("expected status 'ok', got %s", body.Status)
	}

	if body.Uptime == "" {
		t.Error("expected uptime to be non-empty")
	}
}
