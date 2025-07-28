package main

import (
	"log"
	"os"
	"strconv"

	"dougdomingos.com/image-filters/internals/api"
	"dougdomingos.com/image-filters/internals/utils"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	requiredVars := []string{"API_OUTPUT_DIR", "MAX_REQUEST_SIZE"}
	if err := utils.EnsureRequiredEnvVars(requiredVars); err != nil {
		log.Fatal(err)
	}

	serverPort, err := strconv.Atoi(os.Getenv("API_SERVER_PORT"))
	if err != nil {
		serverPort = 8080
	}

	api.StartHTTPServer(serverPort)
}
