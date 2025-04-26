package filters

import (
	"image"
	"math"

	"dougdomingos.com/image-filters/utils"
)

// BinarizationPipeline defines the sequence of steps require to apply the
// binarization filter.
//
// This pipeline includes the following:
//   - Preprocess: applies the Grayscale filter, required to correctly classify
//     the pixels in the image by their intensity.
//   - BuildConcurrent: calculates a global intensity threshold using the whole
//     image, then returns a version of the filter injected with said
//     threshold.
var BinarizationPipeline = FilterPipeline{
	Preprocess:      &GrayscalePipeline,
	Filter:          Binarization,
	BuildConcurrent: buildConcurrentBinarization,
}

var (
	blackPixel = [4]uint8{0, 0, 0, 255}
	whitePixel = [4]uint8{255, 255, 255, 255}
)

// Binarization applies Otsu's thresholding method to binarize the image (i.e.
// classify its pixels based on their intensity). Each pixel is classified as
// black or white based on the obtained threshold:
//
//   - If the intensity of a given pixel is greater than the threshold, it's
//     color is set to while
//   - Otherwise, the pixel's color is set to black
func Binarization(img *image.RGBA, bounds image.Rectangle) {
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

// buildConcurrentBinarization prepares a customized version of the
// Binarization filter for concurrent execution. It calculates a single
// global threshold using the entire image and injects it into a closure,
// ensuring that all goroutines apply a consistent binarization logic.
func buildConcurrentBinarization(img *image.RGBA, bounds image.Rectangle) Filter {
	threshold := otsuThreshold(img, bounds)

	return func(image *image.RGBA, bounds image.Rectangle) {
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
}

// otsuThreshold computes the optimal global threshold for binarization
// based on Otsu's method. It analyzes the intensity histogram of the image
// segment defined by bounds and returns the threshold value that maximizes
// the variance between foreground and background classes.
func otsuThreshold(img *image.RGBA, bounds image.Rectangle) uint8 {
	var (
		histogram        []uint32 = make([]uint32, 256)
		totalPixels      uint32   = uint32(bounds.Dx() * bounds.Dy())
		totalSum         uint64
		sumBackground    float64
		weightBackground float64
		weightForeground float64
		maxVariance      float64
		bestThreshold    uint8
	)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray, _, _, _ := utils.GetRGBA8(img, x, y)
			intensity := gray
			histogram[intensity]++
		}
	}

	for t := range 256 {
		totalSum += uint64(t) * uint64(histogram[t])
	}

	for t := range 256 {
		weightBackground += float64(histogram[t])
		if weightBackground == 0 {
			continue
		}

		weightForeground = float64(totalPixels) - weightBackground
		if weightForeground == 0 {
			break
		}

		sumBackground += float64(t) * float64(histogram[t])

		meanBackground := sumBackground / weightBackground
		meanForeground := (float64(totalSum) - sumBackground) / weightForeground

		varianceBetween := weightBackground * weightForeground * math.Pow(meanBackground-meanForeground, 2)

		if varianceBetween > maxVariance {
			maxVariance = varianceBetween
			bestThreshold = uint8(t)
		}
	}

	return bestThreshold
}
