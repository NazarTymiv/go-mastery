package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/models"
	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/response"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data models.EchoData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if data.Data == "" {
		response.Send(w, response.Error{Error: "missing 'data' field"})
		return
	}

	response.SendJSON(w, http.StatusOK, data)
}
