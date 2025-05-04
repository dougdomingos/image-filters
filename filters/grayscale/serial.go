package grayscale

import (
	"image"

	"dougdomingos.com/image-filters/filters/imgutil"
)

// serialGrayscale applies the grayscale filter to the entire image in a single
// pass. It computes the Rec. 601 luma transform for each pixel in the image,
// updating its color channels to match its result.
func serialGrayscale(img *image.RGBA) {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := (y - img.Rect.Min.Y) * img.Stride
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := rowStart + (x-img.Rect.Min.X)*4

			r, g, b, a := imgutil.GetRGBA8(img, x, y)
			gray := runLumaTransform(r, g, b)

			copy(img.Pix[offset:offset+4], []uint8{gray, gray, gray, a})
		}
	}
}
