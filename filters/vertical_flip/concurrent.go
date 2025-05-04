package vertical_flip

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
)

// concurrentVerticalFlip applies the vertical flip filter to the entire image
// using multiple goroutines. The image is divided into vertical strips to
// ensure the correct mirrored layout.
func concurrentVerticalFlip(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetVerticalPartitions(bounds, numWorkers)
		wg          sync.WaitGroup
	)

	wg.Add(len(imageStrips))
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

	for deltaY := range middle {
		rowStartTop := (bounds.Min.Y + deltaY) * img.Stride
		rowStartBottom := (bounds.Max.Y - 1 - deltaY) * img.Stride

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			topOffset := rowStartTop + (x * 4)
			bottomOffset := rowStartBottom + (x * 4)

			imgutil.SwapPixels(img, topOffset, bottomOffset)
		}
	}
}
