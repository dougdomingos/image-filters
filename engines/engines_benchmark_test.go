package engines_test

import (
	"flag"
	"image"
	"testing"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
)

var (
	// filterName specifies the name of the filter to use during benchmarking.
	filterName = flag.String("filter", "", "Name of the filter to use in benchmarks")

	// imageSize defines the width and height (in pixels) of the square dummy
	// image used for benchmarking.
	imageSize = flag.Int("imageSize", 5000, "Width and height (in pixels) of the square dummy image used for benchmarking.")
)

// BenchmarkExecuteSerial measures the performance and memory allocations of
// applying a filter serially to an image.
//
// It retrieves the specified filter pipeline, prepares a dummy image, and runs
// the serial execution multiple times, reporting time and allocation statistics.
func BenchmarkExecuteSerial(b *testing.B) {
	img := generateDummyImage(*imageSize)
	pipeline, err := filters.GetFilterPipeline(*filterName)
	if err != nil {
		b.Fatalf("Unknown filter: %s", *filterName)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engines.ExecuteSerial(img, pipeline)
	}
}

// BenchmarkExecuteConcurrent measures the performance and memory allocations
// of applying a filter concurrently to an image. It retrieves the specified
// filter pipeline, prepares a dummy image, and runs the concurrent execution
// multiple times, reporting time and allocation statistics.
func BenchmarkExecuteConcurrent(b *testing.B) {
	img := generateDummyImage(*imageSize)
	pipeline, err := filters.GetFilterPipeline(*filterName)
	if err != nil {
		b.Fatalf("Unknown filter: %s", *filterName)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engines.ExecuteConcurrent(img, pipeline)
	}
}

// generateDummyImage creates a new empty RGBA image with the specified size.
// The image is a square of dimensions size x size pixels.
func generateDummyImage(size int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, size, size))
}
