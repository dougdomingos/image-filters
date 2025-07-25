package services

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"strconv"
	"strings"
	"time"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/internals/api/dto"
	"dougdomingos.com/image-filters/pipelines"
)

// BenchmarkHandler provides the benchmark service to the API. It receives the
// filter and a size for the test image through URL parameters.
func BenchmarkHandler(w http.ResponseWriter, r *http.Request) {
	requestData, statusCode, errorMsg := parseBenchmarkRequest(r)
	if statusCode != http.StatusOK {
		http.Error(w, errorMsg, statusCode)
	}

	dummyImg := image.NewRGBA(image.Rect(0, 0, requestData.TestImageSize, requestData.TestImageSize))
	serialRuntime := benchmarkPipeline(requestData.Recipe, *dummyImg)
	concurrentRuntime := benchmarkPipeline(requestData.Recipe, *dummyImg)

	response := dto.BuildBenchmarkResponse(requestData.TestImageSize, serialRuntime, concurrentRuntime)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// parseBenchmarkRequest extract and validates parameters from the HTTP request
// intended for the benchmark service. It ensures the presence of required
// fields, and converts values as required.
func parseBenchmarkRequest(r *http.Request) (*dto.BenchmarkRequestDTO, int, string) {
	filters := r.URL.Query().Get("filters")
	if filters == "" {
		return nil, http.StatusBadRequest, "Parameter \"filters\" is required"
	}

	sampleSize := r.URL.Query().Get("sample-size")
	if sampleSize == "" {
		return nil, http.StatusBadRequest, "Parameter \"sample-size\" is required"
	}
	castedSampleSize, err := strconv.Atoi(sampleSize)
	if err != nil {
		return nil, http.StatusBadRequest, "Parameter \"sample-size\" must be numeric"
	}

	recipe, err := pipelines.BuildRecipe(strings.Split(filters, ","))
	if err != nil {
		return nil, http.StatusNotFound, fmt.Sprintf("Requested filter \"%s\" does not exist", filters)
	}

	return &dto.BenchmarkRequestDTO{
		Recipe:        recipe,
		TestImageSize: castedSampleSize,
	}, http.StatusOK, ""
}

func benchmarkPipeline(recipe pipelines.Recipe, img image.RGBA) int64 {
	start := time.Now()
	engines.ProcessRecipe(&img, &recipe)
	duration := time.Since(start)

	return duration.Milliseconds()
}
