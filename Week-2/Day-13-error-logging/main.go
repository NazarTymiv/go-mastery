package main

import (
	"net/http"

	"github.com/nazartymiv/go-mastery/Week-2/Day-13-error-logging/config"
	"github.com/nazartymiv/go-mastery/Week-2/Day-13-error-logging/logger"
	"github.com/nazartymiv/go-mastery/Week-2/Day-13-error-logging/middleware"
)

func main() {
	config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("pong response", nil)
		w.Write([]byte("pong"))
	})

	handler := middleware.RequestLogger(mux)

	addr := ":8000"
	logger.Info("server starting", map[string]string{"addr": addr})

	if err := http.ListenAndServe(addr, handler); err != nil {
		logger.Error("server failed", err)
	}
}
