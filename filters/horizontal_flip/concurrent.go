package horizontal_flip

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
)

// concurrentHorizontalFlip applies the horizontal flip filter to the entire
// image using multiple goroutines. The image is divided into horizontal strips
// to ensure the correct mirrored layout.
func concurrentHorizontalFlip(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetHorizontalPartitions(bounds, numWorkers)
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
		rowStart := y * img.Stride

		for deltaX := range middle {
			leftOffset := rowStart + (bounds.Min.X+deltaX)*4
			rightOffset := rowStart + (bounds.Max.X-1-deltaX)*4

			imgutil.SwapPixels(img, leftOffset, rightOffset)
		}
	}
}
