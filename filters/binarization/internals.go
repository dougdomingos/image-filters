package binarization

import (
	"image"
	"math"

	"dougdomingos.com/image-filters/filters/imgutil"
)

var (
	// blackPixel represents a fully opaque black pixel in RGBA format
	blackPixel = [4]uint8{0, 0, 0, 255}

	// whitePixel represents a fully opaque white pixel in RGBA format
	whitePixel = [4]uint8{255, 255, 255, 255}
)

// otsuThreshold computes the optimal brightness threshold for the provided
// image. It analyzes the brightness histogram of the image and returns the
// threshold value that maximizes the variance between foreground and
// background components.
func otsuThreshold(img *image.RGBA) uint8 {
	var (
		bounds           image.Rectangle = img.Bounds()
		histogram        []uint32        = make([]uint32, 256)
		totalPixels      uint32          = uint32(bounds.Dx() * bounds.Dy())
		totalSum         uint64
		sumBackground    float64
		weightBackground float64
		weightForeground float64
		maxVariance      float64
		bestThreshold    uint8
	)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			intensity, _, _, _ := imgutil.GetRGBA8(img, x, y)
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
