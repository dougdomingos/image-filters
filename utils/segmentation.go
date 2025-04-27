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
// greatest common divisor (GCD) between the number of CPUs and the segment
// size per  CPU to determine the most efficient number of workers for
// processing.
//
// It calculates the GCD between the number of CPUs and the segment size
// (i.e., the width of the bounds divided by the number of logical CPUs) and
// returns the smaller of the number of CPUs and the GCD.
func GetNumberOfWorkers(bounds image.Rectangle) int {
	numLogicCPUs := runtime.NumCPU()
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}

		return a
	}

	segmentSizePerCPU := bounds.Max.X / numLogicCPUs
	return min(numLogicCPUs, gcd(numLogicCPUs, segmentSizePerCPU))
}
