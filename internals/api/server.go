package api

import (
	"log"
	"net/http"
)

func StartHTTPServer() {
	router := withRequestLogger(buildRouter())

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start HTTP server: %s\n", err.Error())
	}
}
