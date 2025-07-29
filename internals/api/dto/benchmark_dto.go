package dto

import (
	"fmt"
	"math"

	"dougdomingos.com/image-filters/pipelines"
)

// BenchmarkRequestDTO represents the expected parameters for a parsed
// benchmark request.
type BenchmarkRequestDTO struct {

	// Pipeline represents the list of filters to be applied to the dummy image.
	Pipeline pipelines.Pipeline

	// TestImageSize is the requested dimensions for the dummy image to be used
	// on the benchmark tests.
	TestImageSize int
}

// BenchmarkResponseDTO represents the JSON for benchmark HTTP responses.
type BenchmarkResponseDTO struct {

	// ImageSizeInMB is the inferred size (in megabytes), based on the
	// provided dimensions of the test image.
	ImageSizeInMB string `json:"image_size"`

	// ExecTime is the measured execution time for the specified pipeline.
	ExecTime int64 `json:"exec_time"`
}

// BuildBenchmarkResponse constructs a BenchmarkResponseDTO with image size
// formatted to two decimals and runtime metrics in milliseconds.
func BuildBenchmarkResponse(imgDimension int, execTime int64) BenchmarkResponseDTO {
	testImageSize := (math.Pow(float64(imgDimension), 2) * 4) / 1_000_000

	return BenchmarkResponseDTO{
		ImageSizeInMB: fmt.Sprintf("%.4f MB", testImageSize),
		ExecTime:      execTime,
	}
}
