package imgutil

import (
	"image"
	"math"
	"os"
	"runtime"
	"strconv"
)

// GetVerticalSegments splits an image's bounds into an arbitrary number of
// horizontal segments with the same height and, in average, the same width.
func GetVerticalSegments(bounds image.Rectangle, numSegments int) []image.Rectangle {
	segments := make([]image.Rectangle, numSegments)
	avgSegmentWidth := int(math.Ceil(float64(bounds.Max.X) / float64(numSegments)))

	for i := range numSegments {
		startX := bounds.Min.X + (i * avgSegmentWidth)
		endX := min(startX+avgSegmentWidth, bounds.Max.X)

		segments[i] = image.Rect(startX, bounds.Min.Y, endX, bounds.Max.Y)
	}

	return segments
}

// GetHorizontalSegments splits an image's bounds into an arbitrary number of
// horizontal segments with the same width and, in average, the same height.
func GetHorizontalSegments(bounds image.Rectangle, numSegments int) []image.Rectangle {
	segments := make([]image.Rectangle, numSegments)
	avgSegmentHeight := int(math.Ceil(float64(bounds.Max.Y) / float64(numSegments)))

	for i := range numSegments {
		startY := bounds.Min.Y + (i * avgSegmentHeight)
		endY := min(startY+avgSegmentHeight, bounds.Max.Y)

		segments[i] = image.Rect(bounds.Min.X, startY, bounds.Max.X, endY)
	}

	return segments
}

// GetNumberOfWorkers calculates the optimal number of worker goroutines based
// on an arbitrary upper bound (either specified through an environment
// variable or the number of logical CPUs in the system) and the size of the
// image to be processed. It also ensures that, if the average width of the
// segments is less than one, at least one worker should spawn.
func GetNumberOfWorkers(bounds image.Rectangle) int {
	maxWorkers := runtime.NumCPU()

	envVar, isDeclared := os.LookupEnv("MAX_WORKERS_PER_REQUEST")
	if isDeclared && envVar != "" {
		if envWorkersVal, err := strconv.Atoi(envVar); err == nil {
			maxWorkers = envWorkersVal
		}
	}

	segmentWidthPerWorker := bounds.Max.X / maxWorkers
	return max(min(maxWorkers, segmentWidthPerWorker), 1)
}
