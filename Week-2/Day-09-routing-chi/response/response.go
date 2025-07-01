package response

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	Send(w http.ResponseWriter)
}

type Success struct {
	Message string `json:"message"`
}

func (s Success) Send(w http.ResponseWriter) {
	SendJSON(w, http.StatusOK, s)
}

type Error struct {
	Error string `json:"error"`
}

func (e Error) Send(w http.ResponseWriter) {
	SendJSON(w, http.StatusBadRequest, e)
}

func Send(w http.ResponseWriter, r Responder) {
	r.Send(w)
}

func SendJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
