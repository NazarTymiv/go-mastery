package main

import (
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-3/Day-15-middleware/router"
)

func main() {
	log.Printf("Server running on %s", ":8080")

	r := router.SetupRoutes()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err.Error())
	}
}
