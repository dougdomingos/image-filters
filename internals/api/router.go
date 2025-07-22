package api

import (
	"log"
	"net/http"
	"time"
)

func buildRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/process", processorHandler)
	mux.HandleFunc("/bench", benchmarkHandler)

	return mux
}

func withRequestLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        timestamp := time.Now().Format(time.RFC3339)
        log.Printf("[%s] Requested URL: %s", timestamp, r.URL.String())

        next.ServeHTTP(w, r)
    })
}