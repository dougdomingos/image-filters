package vertical_flip

import (
	"image"

	"dougdomingos.com/image-filters/utils"
)

func serialVerticalFlip(img *image.RGBA) {
	bounds := img.Bounds()
	middle := (bounds.Max.Y - bounds.Min.Y) / 2

	for deltaY := range middle {
		rowStartTop := (bounds.Min.Y + deltaY) * img.Stride
		rowStartBottom := (bounds.Max.Y - 1 - deltaY) * img.Stride

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			topOffset := rowStartTop + (x * 4)
			bottomOffset := rowStartBottom + (x * 4)

			utils.SwapPixels(img, topOffset, bottomOffset)
		}
	}
}
