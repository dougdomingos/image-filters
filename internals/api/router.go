package api

import (
	"log"
	"net/http"
	"time"
	"dougdomingos.com/image-filters/internals/api/services"
)

func buildRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/process", services.ProcessorHandler)
	mux.HandleFunc("/bench", services.BenchmarkHandler)

	return mux
}

func withRequestLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        timestamp := time.Now().Format(time.RFC3339)
        log.Printf("[%s] Requested URL: %s", timestamp, r.URL.String())

        next.ServeHTTP(w, r)
    })
}