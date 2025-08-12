// Package vertical_flip implements the vertical flip filter.
//
// The vertical flip filter reverses the order of pixels in each column,
// effectively mirroring the image along its horizontal axis.
package vertical_flip

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

var VerticalFlipFilter = types.NewFilter(nil, VerticalFlip)

// VerticalFlip applies the vertical flip filter to the entire image
// using multiple goroutines. The image is divided into vertical strips to
// ensure the correct mirrored layout.
func VerticalFlip(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetVerticalSegments(bounds, numWorkers)
		wg          sync.WaitGroup
	)

	wg.Add(numWorkers)
	for strip := range imageStrips {
		go verticalFlipWorker(img, imageStrips[strip], &wg)
	}

	wg.Wait()
}

// verticalFlipWorker processes a subregion of the image by reversing the order
// of each column within its respective boundaries.
func verticalFlipWorker(img *image.RGBA, bounds image.Rectangle, wg *sync.WaitGroup) {
	defer wg.Done()
	middle := (bounds.Max.Y - bounds.Min.Y) / 2

	for y := range middle {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			topHead := bounds.Min.Y + y
			bottomHead := bounds.Max.Y - y - 1

			imgutil.SwapPixels(img, x, topHead, x, bottomHead)
		}
	}
}
