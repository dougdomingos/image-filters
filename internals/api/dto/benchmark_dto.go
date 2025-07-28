package dto

import (
	"fmt"
	"math"

	"dougdomingos.com/image-filters/pipelines"
)

// BenchmarkRequestDTO represents the expected parameters for a parsed
// benchmark request.
type BenchmarkRequestDTO struct {

	// Recipe represents the list of filters to be applied to the dummy image.
	Recipe pipelines.Recipe

	// TestImageSize is the requested dimensions for the dummy image to be used
	// on the benchmark tests.
	TestImageSize int
}

// BenchmarkResponseDTO represents the JSON for benchmark HTTP responses.
type BenchmarkResponseDTO struct {

	// TestImageSize is the inferred size (in megabytes), based on the
	// provided dimensions of the test image.
	TestImageSize string `json:"testImageSize"`

	// SerialRuntime is the measured execution time for serial mode.
	SerialRuntime int64 `json:"serialRuntime"`

	// ConcurrentRuntime is the measured execution time for concurrent mode.
	ConcurrentRuntime int64 `json:"concurrentRuntime"`
}

// BuildBenchmarkResponse constructs a BenchmarkResponseDTO with image size
// formatted to two decimals and runtime metrics in milliseconds.
func BuildBenchmarkResponse(imgDimension int, serialTime, concurrentTime int64) BenchmarkResponseDTO {
	testImageSize := (math.Pow(float64(imgDimension), 2) * 4) / 1_000_000

	return BenchmarkResponseDTO{
		TestImageSize:     fmt.Sprintf("%.2f MB", testImageSize),
		SerialRuntime:     serialTime,
		ConcurrentRuntime: concurrentTime,
	}
}
