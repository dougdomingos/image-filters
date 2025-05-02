package binarization

import (
	"image"

	"dougdomingos.com/image-filters/utils"
)

func serialBinarization(img *image.RGBA) {
	bounds := img.Bounds()
	threshold := otsuThreshold(img, bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := (y - img.Rect.Min.Y) * img.Stride
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := rowStart + (x-img.Rect.Min.X)*4
			intensity, _, _, _ := utils.GetRGBA8(img, x, y)

			if intensity > threshold {
				copy(img.Pix[offset:offset+4], whitePixel[:])
			} else {
				copy(img.Pix[offset:offset+4], blackPixel[:])
			}
		}
	}
}
