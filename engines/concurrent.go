package engines

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/utils"
)

// ExecuteConcurrent applies the provided filter to the specified image object
// by splitting it into several segments, assigning each one to a goroutine.
func ExecuteConcurrent(img *image.RGBA, filter filters.Filter, workers int) {
	var wg sync.WaitGroup
	imageStrips := utils.GetVerticalStrips(img.Bounds(), workers)
	
	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go func(s image.Rectangle) {
			defer wg.Done()
			filter(img, s)
		}(imageStrips[strip])
	}

	wg.Wait()
}
