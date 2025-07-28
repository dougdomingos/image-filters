package api

import (
	"fmt"
	"log"
	"net/http"
)

// StartHTTPServer initiates the API HTTP server with the specified routes.
// By default, the server listens to port 8080.
func StartHTTPServer(port int) {
	router := withRequestLogger(buildRouter())
	address := fmt.Sprintf(":%d", port)

	log.Printf("Starting HTTP server on port %d\n", port)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("Could not start HTTP server: %s\n", err.Error())
	}
}
