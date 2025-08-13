// Package horizontal_flip implements an image filter that reverses images in
// the horizontal axis.
package horizontal_flip

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

// HorizontalFlipFilter applies the horizontal flip transformation to a image.
// No pre-processing is performed.
var HorizontalFlipFilter = types.NewFilter(nil, HorizontalFlip)

// HorizontalFlip applies the horizontal flip filter to the entire image. The
// image is divided into horizontal segments, each delegated to a worker
// goroutine.
func HorizontalFlip(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetHorizontalSegments(bounds, numWorkers)
		wg          sync.WaitGroup
	)

	wg.Add(numWorkers)
	for strip := range imageStrips {
		go horizontalFlipWorker(img, imageStrips[strip], &wg)
	}

	wg.Wait()
}

// horizontalFlipWorker reverses the order of pixels for each row in its
// segment.
func horizontalFlipWorker(img *image.RGBA, bounds image.Rectangle, wg *sync.WaitGroup) {
	defer wg.Done()
	middle := (bounds.Max.X - bounds.Min.X) / 2

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := range middle {
			leftHead := bounds.Min.X + x
			rightHead := bounds.Max.X - x - 1

			imgutil.SwapPixels(img, leftHead, y, rightHead, y)
		}
	}
}
