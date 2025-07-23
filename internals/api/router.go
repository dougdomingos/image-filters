package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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
		route := r.URL.Path

		queryParams := r.URL.Query()
		var paramList []string
		for key, values := range queryParams {
			for _, value := range values {
				paramList = append(paramList, fmt.Sprintf("%s: %s", key, value))
			}
		}

		paramStr := ""
		if len(paramList) > 0 {
			paramStr = fmt.Sprintf(" [%s]", strings.Join(paramList, ", "))
		}

		log.Printf("[%s] Request to %s%s", timestamp, route, paramStr)
		next.ServeHTTP(w, r)
    })
}