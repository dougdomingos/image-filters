package services

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"strconv"
	"time"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters/types"
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
	serialRuntime := benchmarkPipeline(requestData.FilterPipeline, *dummyImg, false)
	concurrentRuntime := benchmarkPipeline(requestData.FilterPipeline, *dummyImg, true)

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
	filterName := r.URL.Query().Get("filter")
	if filterName == "" {
		return nil, http.StatusBadRequest, "Parameter \"filter\" is required"
	}

	sampleSize := r.URL.Query().Get("sample-size")
	if sampleSize == "" {
		return nil, http.StatusBadRequest, "Parameter \"sample-size\" is required"
	}
	castedSampleSize, err := strconv.Atoi(sampleSize)
	if err != nil {
		return nil, http.StatusBadRequest, "Parameter \"sample-size\" must be numeric"
	}

	pipeline, err := pipelines.GetFilterPipeline(filterName)
	if err != nil {
		return nil, http.StatusNotFound, fmt.Sprintf("Requested filter \"%s\" does not exist", filterName)
	}

	return &dto.BenchmarkRequestDTO{
		FilterPipeline: pipeline,
		TestImageSize:  castedSampleSize,
	}, http.StatusOK, ""
}

// benchmarkPipeline applies a filter pipeline to an image and returns the
// processing duration in milliseconds. The caller may specify if the pipeline
// should run concurrently or not.
func benchmarkPipeline(pipeline types.FilterPipeline, img image.RGBA, concurrentMode bool) int64 {
	start := time.Now()
	engines.ApplyFilterPipeline(&img, &pipeline, concurrentMode)
	duration := time.Since(start)

	return duration.Milliseconds()
}
