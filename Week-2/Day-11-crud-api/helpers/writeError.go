package helpers

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
