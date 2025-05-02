package grayscale

import (
	"image"

	"dougdomingos.com/image-filters/utils"
)

func serialGrayscale(img *image.RGBA) {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := (y - img.Rect.Min.Y) * img.Stride
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := rowStart + (x-img.Rect.Min.X)*4

			r, g, b, a := utils.GetRGBA8(img, x, y)
			gray := RunLumaTransform(r, g, b)

			copy(img.Pix[offset:offset+4], []uint8{gray, gray, gray, a})
		}
	}
}
