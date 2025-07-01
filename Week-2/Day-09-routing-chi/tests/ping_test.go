package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/handlers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/response"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	handlers.PingHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

	var body response.Success
	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if body.Message != "pong" {
		t.Errorf("expected message 'pong', got '%s'", body.Message)
	}
}
