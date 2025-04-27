package filters

import (
	"image"

	"dougdomingos.com/image-filters/utils"
)

// GrayscalePipeline defines the grayscale filter pipeline.
//
// This pipeline includes the following:
//   - Preprocess: none
//   - BuildConcurrent: none; works on serial and concurrent modes by
//     default
var GrayscalePipeline = BuildFilterPipeline(Grayscale, nil, nil, nil)

// Grayscale converts each pixel within a specified segment of the image into
// its corresponding shade of gray.
//
// This implementation uses the Rec. 601 luma transform to compute the
// grayscale value. The alpha channel is left unmodified.
func Grayscale(img *image.RGBA, bounds image.Rectangle) {
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := (y - img.Rect.Min.Y) * img.Stride
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := rowStart + (x-img.Rect.Min.X)*4

			r, g, b, a := utils.GetRGBA8(img, x, y)
			gray := runLumaTransform(r, g, b)

			copy(img.Pix[offset:offset+4], []uint8{gray, gray, gray, a})
		}
	}
}

// runLumaTransform computes the grayscale value of a pixel based on the Rec. 601
// luma transform:
//
//	gray = (0.299 * R) + (0.587 * G) + (0.114 * B)
//
// The Rec. 601 standard accounts for human visual perception, assigning higher
// weight to the green channel,followed by red and blue, based on the eye's
// sensitivity to each color.
func runLumaTransform(r, g, b uint8) uint8 {
	weightedRed := 0.299 * float64(r)
	weightedGreen := 0.587 * float64(g)
	weightedBlue := 0.114 * float64(b)

	return uint8(weightedRed + weightedGreen + weightedBlue)
}
