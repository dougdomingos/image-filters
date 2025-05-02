package vertical_flip

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/partitions"
	"dougdomingos.com/image-filters/utils"
)

func concurrentVerticalFlip(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = 8
		imageStrips = partitions.GetVerticalPartitions(bounds, numWorkers)
		wg          sync.WaitGroup
	)

	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go VerticalFlipWorker(img, imageStrips[strip], &wg)
	}

	wg.Wait()
}

func VerticalFlipWorker(img *image.RGBA, bounds image.Rectangle, wg *sync.WaitGroup) {
	defer wg.Done()
	middle := (bounds.Max.Y - bounds.Min.Y) / 2

	for deltaY := range middle {
		rowStartTop := (bounds.Min.Y + deltaY) * img.Stride
		rowStartBottom := (bounds.Max.Y - 1 - deltaY) * img.Stride

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			topOffset := rowStartTop + (x * 4)
			bottomOffset := rowStartBottom + (x * 4)

			utils.SwapPixels(img, topOffset, bottomOffset)
		}
	}
}
