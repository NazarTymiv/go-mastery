package handlers

import (
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/response"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	response.Send(w, response.Success{Message: "pong"})
}
