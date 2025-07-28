package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"dougdomingos.com/image-filters/internals/api/services"
)

// buildRouter creates and configures a ServeMux with the handlers for each
// route avaliable on the API.
func buildRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/process", services.ProcessorHandler)
	mux.HandleFunc("/bench", services.BenchmarkHandler)
	mux.HandleFunc("/images/", createFileHandler(os.Getenv("API_OUTPUT_DIR"), "/images/"))

	return mux
}

// withRequestLogger acts as a middleware handler that captures requests to the
// API and logs basic information for each one (e.g., requested route, URL
// params).
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

// createFileHandler creates a http.HandlerFunc specialized on serving requests
// to images within the specified output directory.
func createFileHandler(outputDir string, prefix string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(outputDir))

	return func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(prefix, fs).ServeHTTP(w, r)
	}
}
