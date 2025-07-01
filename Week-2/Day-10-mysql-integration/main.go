package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/config"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/db"
	"github.com/nazartymiv/go-mastery/Week-2/Day-10-mysql-integration/router"
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
