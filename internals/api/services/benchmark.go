package services

import (
	"log"
	"net/http"
)

func BenchmarkHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Benchmark route accessed")
}
