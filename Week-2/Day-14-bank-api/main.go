package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/config"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/db"
	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/router"
)

func main() {
	cfg := config.Load()
	database := db.Connect(cfg.DBDSN)
	defer database.Close()

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Server running on %s", addr)

	r := router.SetupRoutes(database)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
