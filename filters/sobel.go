package filters

import (
	"image"
	"math"

	"dougdomingos.com/image-filters/utils"
)

// SobelPipeline defines the Sobel's edge-detection filter pipeline.
//
// This pipeline includes the following:
//   - Preprocess: none
//   - BuildConcurrent: creates a full copy of the original image, returning a
//     version of the filter injected with said copy
var SobelPipeline = BuildFilterPipeline(Sobel, nil, buildConcurrentSobel, nil)

// SobelGrayscaledPipeline defines the Sobel's edge detection filter pipeline
// with a grayscale pre-processing step.
//
// This pipeline includes the following:
//   - Preprocess: applies the Grayscale filter, which enhances edge detection
//     by considering the intensity of the pixel's colors instead of color
//     variation.
//   - BuildConcurrent: creates a full copy of the original image, returning a
//     version of the filter injected with said copy
var SobelGrayscaledPipeline = BuildFilterPipeline(Sobel, &GrayscalePipeline, buildConcurrentSobel, nil)

var (
	// Sobel's convolution kernel used for computing horizontal gradients.
	gX = [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	// Sobel's convolution kernel used for computing vertical gradients.
	gY = [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}
)

// Sobel applies the Sobel filter to the given image in-place over the
// specified bounds. It calculates gradient magnitude for each channel
// separately, resulting in a color-preserving edge map.
func Sobel(img *image.RGBA, bounds image.Rectangle) {
	originalImg := copyImagePartition(img, bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := (y - img.Rect.Min.Y) * img.Stride
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := rowStart + (x-img.Rect.Min.X)*4

			var r8, g8, b8, a8 uint8
			var gxR, gxG, gxB, gyR, gyG, gyB int

			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					deltaX := x + kx
					deltaY := y + ky

					// If the current neighbor is outside of the image, jump to next iteration
					if !isPositionWithinImage(img, deltaX, deltaY) {
						continue
					}

					r8, g8, b8, a8 = utils.GetRGBA8(&originalImg, deltaX, deltaY)
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

// buildConcurrentSobel returns a concurrent-safe Sobel filter function to be
// used within concurrent pipelines. It operates on a global copy of the source
// image to avoid race conditions during concurrent execution.
func buildConcurrentSobel(img *image.RGBA, bounds image.Rectangle) Filter {
	originalImg := copyImagePartition(img, bounds)

	return func(image *image.RGBA, bounds image.Rectangle) {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			rowStart := (y - img.Rect.Min.Y) * img.Stride
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				offset := rowStart + (x-img.Rect.Min.X)*4

				var r8, g8, b8, a8 uint8
				var gxR, gxG, gxB, gyR, gyG, gyB int

				for ky := -1; ky <= 1; ky++ {
					for kx := -1; kx <= 1; kx++ {
						deltaX := x + kx
						deltaY := y + ky

						// If the current neighbor is outside of the image, jump to next iteration
						if !isPositionWithinImage(img, deltaX, deltaY) {
							continue
						}

						r8, g8, b8, a8 = utils.GetRGBA8(&originalImg, deltaX, deltaY)
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
}

// clampColorValue ensures that the provided value stays within the range of an
// RGBA color channel value (i.e., from 0 to 255).
func clampColorValue(val float64) uint8 {
	if val < 0 {
		return 0
	}

	if val > 255 {
		return 255
	}

	return uint8(val)
}

// copyImagePartition creates a copy of a region of the provided image.
func copyImagePartition(img *image.RGBA, bounds image.Rectangle) image.RGBA {
	copiedPartition := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			srcOffset := img.PixOffset(x, y)
			dstOffset := copiedPartition.PixOffset(x, y)

			copy(copiedPartition.Pix[dstOffset:dstOffset+4], img.Pix[srcOffset:srcOffset+4])
		}
	}

	return *copiedPartition
}

// isPositionWithinImage checks if the pixel at the specified coordinates is
// within the boundaries of the provided image.
func isPositionWithinImage(img *image.RGBA, x, y int) bool {
	isWithinXBounds := (x >= img.Rect.Min.X) && (x < img.Rect.Max.X)
	isWithinYBounds := (y >= img.Rect.Min.Y) && (y < img.Rect.Max.Y)

	return isWithinXBounds && isWithinYBounds
}
