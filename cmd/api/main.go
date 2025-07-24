package main

import (
	"log"

	"dougdomingos.com/image-filters/internals/api"
	"dougdomingos.com/image-filters/internals/utils"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	requiredEnvVars := []string{"API_OUTPUT_DIR", "MAX_REQUEST_SIZE"}
	for _, key := range requiredEnvVars {
		if x, err := utils.GetEnvVar(key); err != nil {
			log.Fatalf("[ERROR] Missing required environment variable: %s\n", key)
		} else {
			log.Print(key, x)
		}
	}

	api.StartHTTPServer()
}
