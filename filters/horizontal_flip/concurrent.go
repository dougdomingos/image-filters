package horizontal_flip

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/partitions"
	"dougdomingos.com/image-filters/utils"
)

func concurrentHorizontalFlip(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = 8
		imageStrips = partitions.GetHorizontalPartitions(bounds, numWorkers)
		wg          sync.WaitGroup
	)

	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go HorizontalFlipWorker(img, imageStrips[strip], &wg)
	}

	wg.Wait()
}

func HorizontalFlipWorker(img *image.RGBA, bounds image.Rectangle, wg *sync.WaitGroup) {
	defer wg.Done()
	middle := (bounds.Max.X - bounds.Min.X) / 2

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := y * img.Stride

		for deltaX := range middle {
			leftOffset := rowStart + (bounds.Min.X+deltaX)*4
			rightOffset := rowStart + (bounds.Max.X-1-deltaX)*4

			utils.SwapPixels(img, leftOffset, rightOffset)
		}
	}
}
