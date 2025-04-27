package filters

import (
	"image"

	"dougdomingos.com/image-filters/utils"
)

// HorizontalFlipPipeline defines the horizontal flip filter pipeline.
//
// This pipeline includes the following:
//   - Preprocess: none
//   - BuildConcurrent: none; works on serial and concurrent modes.
//     Requires horizontal segmentation for concurrent mode.
var HorizontalFlipPipeline = FilterPipeline{
	Filter: HorizontalFlip,
}

// HorizontalFlip reverses the order of each line of pixels within the
// specified segment of the image.
func HorizontalFlip(img *image.RGBA, bounds image.Rectangle) {
	middle := (bounds.Max.X - bounds.Min.X) / 2

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := y * img.Stride

		for deltaX := range middle {
			leftOffset := rowStart + (bounds.Min.X + deltaX) * 4
			rightOffset := rowStart + (bounds.Max.X - 1 - deltaX) * 4

			utils.SwapPixels(img, leftOffset, rightOffset)
		}
	}
}
