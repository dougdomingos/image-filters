// Package vertical_flip implements an image filter that reverses images in
// the vertical axis.
package vertical_flip

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

// VerticalFlipFilter applies the vertical flip transformation to a image. No
// pre-process is performed.
var VerticalFlipFilter = types.NewFilter(nil, VerticalFlip)

// VerticalFlip applies the vertical flip filter to the entire image. The image
// is divided into vertical segments, each delegated to a worker goroutine.
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

// verticalFlipWorker reverses the order of pixels for each column in its
// segment.
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
