package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func RequestTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("\nRequest took: %dms\n", time.Since(start).Milliseconds())
	})
}
