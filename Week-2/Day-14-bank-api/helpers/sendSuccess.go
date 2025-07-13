package helpers

import (
	"encoding/json"
	"net/http"
)

func SendSuccess(w http.ResponseWriter, data interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
