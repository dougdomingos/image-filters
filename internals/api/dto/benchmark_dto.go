package dto

import (
	"fmt"
	"math"

	"dougdomingos.com/image-filters/filters/types"
)

type BenchmarkRequestDTO struct {
	FilterPipeline types.FilterPipeline
	TestImageSize int
}

type BenchmarkResponseDTO struct {
	TestImageSize     string `json:"testImageSize"`
	SerialRuntime     int64  `json:"serialRuntime"`
	ConcurrentRuntime int64  `json:"concurrentRuntime"`
}

func BuildBenchmarkResponse(imgDimension int, serialTime, concurrentTime int64) BenchmarkResponseDTO {
	testImageSize := (math.Pow(float64(imgDimension), 2) * 4) / 1_000_000

	return BenchmarkResponseDTO{
		TestImageSize: fmt.Sprintf("%.2f MB", testImageSize),
		SerialRuntime: serialTime,
		ConcurrentRuntime: concurrentTime,
	}
}