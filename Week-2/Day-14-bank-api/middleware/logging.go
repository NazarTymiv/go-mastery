package middleware

import (
	"net/http"
	"time"

	"github.com/nazartymiv/go-mastery/Week-2/Day-14-bank-api/logger"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rr := &responseRecorder{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(rr, r)

		duration := time.Since(start)
		logger.Info("request completed", map[string]interface{}{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   rr.statusCode,
			"duration": duration.String(),
		})
	})
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}
