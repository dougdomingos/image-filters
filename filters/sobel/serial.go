package sobel

import (
	"image"
	"math"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/utils"
)

// serialSobel applies the sobel filter to the entire image in a single pass.
// It creates a full copy of the original image, using it to calculate the
// color channel gradients for each pixel of the image.
func serialSobel(img *image.RGBA) {
	bounds := img.Bounds()
	copyImg := imgutil.CopyPaddedImagePartition(img, bounds, copyPadding)

	paddedMinX, paddedMaxX := copyImg.Rect.Min.X+copyPadding, copyImg.Rect.Max.X-copyPadding
	paddedMinY, paddedMaxY := copyImg.Rect.Min.Y+copyPadding, copyImg.Rect.Max.Y-copyPadding

	for y := paddedMinY; y < paddedMaxY; y++ {
		srcRowStart := (y - copyPadding) * img.Stride
		for x := paddedMinX; x < paddedMaxX; x++ {
			offset := srcRowStart + (x-copyPadding)*4

			var r8, g8, b8, a8 uint8
			var gxR, gxG, gxB, gyR, gyG, gyB int

			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					deltaX := x + kx
					deltaY := y + ky

					r8, g8, b8, a8 = utils.GetRGBA8(&copyImg, deltaX, deltaY)
					r := int(r8)
					g := int(g8)
					b := int(b8)

					valKernelX := gX[ky+1][kx+1]
					valKernelY := gY[ky+1][kx+1]

					gxR += int(valKernelX * r)
					gxG += int(valKernelX * g)
					gxB += int(valKernelX * b)

					gyR += int(valKernelY * r)
					gyG += int(valKernelY * g)
					gyB += int(valKernelY * b)
				}
			}

			gradR := clampColorValue(math.Sqrt(float64(gxR*gxR + gyR*gyR)))
			gradG := clampColorValue(math.Sqrt(float64(gxG*gxG + gyG*gyG)))
			gradB := clampColorValue(math.Sqrt(float64(gxB*gxB + gyB*gyB)))

			copy(img.Pix[offset:offset+4], []uint8{gradR, gradG, gradB, a8})
		}
	}
}
