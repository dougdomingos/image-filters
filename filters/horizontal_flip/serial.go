package horizontal_flip

import (
	"image"

	"dougdomingos.com/image-filters/utils"
)

// serialHorizontalFlip applies the horizontal flip filter to the entire image
// in a single pass. It reverses the order of the pixels on each row.
func serialHorizontalFlip(img *image.RGBA) {
	bounds := img.Bounds()
	middle := (bounds.Max.X - bounds.Min.X) / 2

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := y * img.Stride

		for deltaX := range middle {
			leftOffset := rowStart + (bounds.Min.X+deltaX)*4
			rightOffset := rowStart + (bounds.Max.X-1-deltaX)*4

			utils.SwapPixels(img, leftOffset, rightOffset)
		}
	}
}
