package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/handlers"
	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/models"
)

func TestEchoHandler_WithData(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewBufferString(`{"data": "asd"}`))
	w := httptest.NewRecorder()

	handlers.EchoHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}

	var body models.EchoData
	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if body.Data != "asd" {
		t.Errorf("expected body data 'asd', got %s", body.Data)
	}
}

func TestEchoHandler_WithoutData(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/echo", nil)
	w := httptest.NewRecorder()

	handlers.EchoHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", res.StatusCode)
	}
}
