package partitions

import (
	"image"
	"math"
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