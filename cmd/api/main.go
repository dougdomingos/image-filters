package main

import (
	"dougdomingos.com/image-filters/internals/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	api.StartHTTPServer()
}
