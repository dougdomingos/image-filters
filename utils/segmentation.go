package utils

import (
	"image"
	"math"
	"runtime"
)

// GetVerticalStrips splits an image's bounds into an arbitrary number of
// vertical strips, with the same height and, roughly, the same width.
func GetVerticalStrips(bounds image.Rectangle, stripCount int) []image.Rectangle {
	strips := make([]image.Rectangle, stripCount)
	imageWidth := bounds.Dx()
	stripWidth := int(math.Ceil(float64(imageWidth) / float64(stripCount)))

	for i := range stripCount {
		startX := bounds.Min.X + (i * stripWidth)
		endX := min(startX+stripWidth, imageWidth)

		strips[i] = image.Rect(startX, bounds.Min.Y, endX, bounds.Max.Y)
	}

	return strips
}

// GetNumberOfWorkers calculates the optimal number of workers based on the
// number of logical CPUs available and the size of the task (given by the
// bounds). The function considers both the number of available CPUs and the
// width of each worker's segment. It also ensures that, if the width of the
// segments is less than zero, the returned value would be 1.
func GetNumberOfWorkers(bounds image.Rectangle) int {
	numLogicCPUs := runtime.NumCPU()
	segmentWidthPerLogicCPU := bounds.Max.X / numLogicCPUs

	return max(min(numLogicCPUs, segmentWidthPerLogicCPU), 1)
}
