// Package binarization implements an image filter that classifies each pixel
// as foreground or background based on a global brightness threshold. Pixels
// whose brightness is higher than the threshold are marked as white, while
// those below are marked as black.
//
// This implementation uses Otsu's Method to compute thresholds of images, as
// it adapts to the brightness distribution of the image and computes a
// threshold that best separates foreground and background pixels, ensuring
// consistent results regardless of lighting variations.
//
// As binarization depends on pixel brightness, this implementation declares
// the [Grayscale] filter as a pre-processing step to improve threshold
// computations.
package binarization

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

// BinarizationFilter applies a grayscale pre-process followed by the
// binarization transformation.
var BinarizationFilter = types.NewFilter(grayscale.Grayscale, Binarization)

// Binarization applies the binarization transformation to the entire image. The
// image is divided into vertical segments, each delegated to a worker goroutine.
func Binarization(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetVerticalSegments(bounds, numWorkers)
		threshold   = otsuThreshold(img)
		wg          sync.WaitGroup
	)

	wg.Add(numWorkers)
	for strip := range imageStrips {
		go binarizationWorker(img, imageStrips[strip], threshold, &wg)
	}

	wg.Wait()
}

// binarizationWorker processes a segment of the image based on the global
// brightness threshold. Pixels whose brightness is above the threshold are
// set to while, while others are set to black.
func binarizationWorker(img *image.RGBA, segment image.Rectangle, threshold uint8, wg *sync.WaitGroup) {
	defer wg.Done()

	for y := segment.Min.Y; y < segment.Max.Y; y++ {
		for x := segment.Min.X; x < segment.Max.X; x++ {
			intensity, _, _, _ := imgutil.GetRGBA8(img, x, y)

			if intensity > threshold {
				imgutil.WriteToPixel(img, x, y, whitePixel)
			} else {
				imgutil.WriteToPixel(img, x, y, blackPixel)
			}

		}
	}
}
