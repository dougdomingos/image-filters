package engines_test

import (
	"flag"
	"image"
	"os"
	"strings"
	"testing"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/pipelines"
)

var (
	// filterName specifies the name of the filter to use during benchmarking.
	filterName = flag.String("filters", "", "Name of the filter to use in benchmarks")

	// imageSize defines the width and height (in pixels) of the square dummy
	// image used for benchmarking.
	imageSize = flag.Int("imageSize", 5000, "Width and height (in pixels) of the square dummy image used for benchmarking.")

	// sinkPixel servers as a anchor value to ensure that the Go's compiler
	// won't skip the filter's execution. It is used to store the first value
	// in the image.Pix[] array.
	sinkPixel uint8
)

// BenchmarkExecuteSerial measures the performance and memory allocations of
// applying a filter serially to an image.
//
// It retrieves the specified filter pipeline, prepares a dummy image, and runs
// the serial execution multiple times, reporting time and allocation statistics.
func BenchmarkExecuteSerial(b *testing.B) {
	os.Setenv("MAX_WORKERS_PER_REQUEST", "1")
	recipe, err := pipelines.NewPipeline(strings.Split(*filterName, ","))
	if err != nil {
		b.Fatalf("Unknown filter: %s", *filterName)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		img := generateDummyImage(*imageSize)
		engines.ProcessPipeline(img, &recipe)
		sinkPixel = img.Pix[0]
	}
}

// BenchmarkExecuteConcurrent measures the performance and memory allocations
// of applying a filter concurrently to an image. It retrieves the specified
// filter pipeline, prepares a dummy image, and runs the concurrent execution
// multiple times, reporting time and allocation statistics.
func BenchmarkExecuteConcurrent(b *testing.B) {
	os.Unsetenv("MAX_WORKERS_PER_REQUEST")
	recipe, err := pipelines.NewPipeline(strings.Split(*filterName, ","))
	if err != nil {
		b.Fatalf("Unknown filter: %s", *filterName)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		img := generateDummyImage(*imageSize)
		engines.ProcessPipeline(img, &recipe)
		sinkPixel = img.Pix[0]
	}
}

// generateDummyImage creates a new empty RGBA image with the specified size.
// The image is a square of dimensions size x size pixels.
func generateDummyImage(size int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, size, size))
}
