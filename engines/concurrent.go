package engines

import (
	"image"
	"runtime"
	"sync"

	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/utils"
)

// ExecuteConcurrent applies the provided filter to the specified image object
// by splitting it into several segments, assigning each one to a goroutine.
func ExecuteConcurrent(img *image.RGBA, filter filters.Filter) {
	var (
		workers     = runtime.NumCPU()
		imageStrips = utils.GetVerticalStrips(img.Bounds(), workers)
		wg sync.WaitGroup
	)
	
	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go func(s image.Rectangle) {
			defer wg.Done()
			filter(img, s)
		}(imageStrips[strip])
	}

	wg.Wait()
}
