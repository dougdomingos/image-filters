package api

import (
	"log"
	"net/http"
)

// StartHTTPServer initiates the API HTTP server with the specified routes.
// By default, the server listens to port 8080.
func StartHTTPServer() {
	router := withRequestLogger(buildRouter())

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start HTTP server: %s\n", err.Error())
	}
}
