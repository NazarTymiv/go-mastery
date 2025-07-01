package handlers

import (
	"fmt"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/response"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		response.Send(w, response.Error{Error: "Missing name"})
		return
	}

	message := fmt.Sprintf("Hello, %s", name)
	response.Send(w, response.Success{Message: message})
}
