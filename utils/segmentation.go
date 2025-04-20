package utils

import (
	"image"
	"math"
)

// GetVerticalStrips splits an image's bounds into an arbitrary number of
// vertical strips, with the same height and, roughly, the same width.
func GetVerticalStrips(bounds image.Rectangle, stripCount int) []image.Rectangle {
	strips := make([]image.Rectangle, stripCount)
	imageWidth := bounds.Dx()
	stripWidth := int(math.Ceil(float64(imageWidth) / float64(stripCount)))

	for i := range stripCount {
		startX := bounds.Min.X + (i * stripWidth)
		endX   := min(startX + stripWidth, imageWidth)
		
		strips[i] = image.Rect(startX, bounds.Min.Y, endX, bounds.Max.Y)
	}

	return strips
}
