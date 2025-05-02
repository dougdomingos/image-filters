package partitions

import (
	"image"
	"math"
	"runtime"
)

// GetVerticalPartitions splits an image's bounds into an arbitrary number of
// horizontal partitions with the same height and, in average, the same width.
func GetVerticalPartitions(bounds image.Rectangle, segments int) []image.Rectangle {
	partitions := make([]image.Rectangle, segments)
	partitionWidth := int(math.Ceil(float64(bounds.Max.X) / float64(segments)))

	for i := range segments {
		startX := bounds.Min.X + (i * partitionWidth)
		endX := min(startX + partitionWidth, bounds.Max.X)

		partitions[i] = image.Rect(startX, bounds.Min.Y, endX, bounds.Max.Y)
	}

	return partitions
}

// GetHorizontalPartitions splits an image's bounds into an arbitrary number of
// horizontal partitions with the same width and, in average, the same height.
func GetHorizontalPartitions(bounds image.Rectangle, segments int) []image.Rectangle {
	partitions := make([]image.Rectangle, segments)
	partitionHeight := int(math.Ceil(float64(bounds.Max.Y) / float64(segments)))

	for i := range segments {
		startY := bounds.Min.Y + (i * partitionHeight)
		endY := min(startY + partitionHeight, bounds.Max.Y)

		partitions[i] = image.Rect(bounds.Min.X, startY, bounds.Max.X, endY)
	}

	return partitions
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