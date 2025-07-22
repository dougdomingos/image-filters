package api

import (
	"fmt"
	"net/http"
)

func benchmarkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Benchmark route accessed")
}

func processorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Image processor route accessed")
}
