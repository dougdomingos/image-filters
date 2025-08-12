// Package binarization implements the binarization filter.
//
// The binarization filter converts an image to black and white based on a
// brightness threshold, which can either be fixed or calculated based on the
// image's pixels.
//
// This implementation applies Otsu's Method of thresholding to determine the
// value that best separates the foreground and background components of the
// image.
package binarization

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

var BinarizationFilter = types.NewFilter(grayscale.Grayscale, Binarization)

// Binarization applies the binarization filter to the entire image
// using multiple goroutines. It first computes Otsu's global threshold,
// then partitions the image and processes each partition concurrently.
func Binarization(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetVerticalSegments(bounds, numWorkers)
		threshold   = otsuThreshold(img, bounds)
		wg          sync.WaitGroup
	)

	wg.Add(numWorkers)
	for strip := range imageStrips {
		go binarizationWorker(img, imageStrips[strip], threshold, &wg)
	}

	wg.Wait()
}

// binarizationWorker processes a subregion of the image by applying the
// binarization filter based on a shared global threshold. It updates each
// pixel in the subregion to either black or white, depending on it's
// intensity.
func binarizationWorker(img *image.RGBA, bounds image.Rectangle, threshold uint8, wg *sync.WaitGroup) {
	defer wg.Done()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			intensity, _, _, _ := imgutil.GetRGBA8(img, x, y)

			if intensity > threshold {
				imgutil.WriteToPixel(img, x, y, whitePixel)
			} else {
				imgutil.WriteToPixel(img, x, y, blackPixel)
			}

		}
	}
}
