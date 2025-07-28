package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"dougdomingos.com/image-filters/internals/utils"
)

// StartHTTPServer initiates the API HTTP server with the specified routes.
// By default, the server listens to port 8080.
func StartHTTPServer() {
	router := withRequestLogger(buildRouter())
	port := getServerPort()
	address := fmt.Sprintf(":%d", port)

	requiredVars := []string{"API_OUTPUT_DIR", "MAX_REQUEST_SIZE"}
	if err := utils.EnsureRequiredEnvVars(requiredVars); err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting HTTP server on port %d\n", port)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("Could not start HTTP server: %s\n", err.Error())
	}
}

func getServerPort() int {
	serverPort, err := strconv.Atoi(os.Getenv("API_SERVER_PORT"))
	if err != nil {
		serverPort = 8080
	}

	return serverPort
}
