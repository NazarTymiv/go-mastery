package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/handlers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/response"
)

func TestGreetHandler_WithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet?name=Nazar", nil)
	w := httptest.NewRecorder()

	handlers.GreetHandler(w, req)

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

	if body.Message != "Hello, Nazar" {
		t.Errorf("expected message 'Hello, Nazar', got %s", body.Message)
	}
}

func TestGreetHandler_WithoutName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	w := httptest.NewRecorder()

	handlers.GreetHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status code 400, got %d", res.StatusCode)
	}

	var body response.Error
	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if body.Error != "Missing name" {
		t.Errorf("expected message 'Missing name', got %s", body.Error)
	}
}
