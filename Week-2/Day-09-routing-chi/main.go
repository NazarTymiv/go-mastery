package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nazartymiv/go-mastery/Week-2/Day-09-routing-chi/router"
)

var serverStartTime time.Time

func main() {
	serverStartTime = time.Now()

	r := router.SetupRoutes(serverStartTime)

	fmt.Println("Server started at http://localhost:8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		fmt.Println("Error running server:", err)
	}
}
