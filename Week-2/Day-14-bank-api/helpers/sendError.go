package helpers

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
