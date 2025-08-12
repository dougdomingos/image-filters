package horizontal_flip

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
)

// HorizontalFlip applies the horizontal flip filter to the entire
// image using multiple goroutines. The image is divided into horizontal strips
// to ensure the correct mirrored layout.
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

// horizontalFlipWorker processes a subregion of the image by reversing the
// order of each row within its respective boundaries.
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
