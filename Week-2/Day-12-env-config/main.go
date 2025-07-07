package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-12-env-config/config"
)

func main() {
	config.Load()

	fmt.Printf("Running in %s mode\n", config.AppConfig.Env)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"Port":  config.AppConfig.Port,
			"DBDSN": config.AppConfig.DBDSN,
			"Env":   config.AppConfig.Env,
		})
	})

	addr := fmt.Sprintf(":%d", config.AppConfig.Port)
	log.Printf("Server listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
