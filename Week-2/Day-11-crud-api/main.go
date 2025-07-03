package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/config"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/db"
	"github.com/nazartymiv/go-mastery/Week-2/Day-11-crud-api/router"
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
