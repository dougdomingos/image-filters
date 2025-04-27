package engines

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/utils"
)

// ExecuteConcurrent applies the provided filter to the specified image object
// by splitting it into several segments, assigning each one to a goroutine.
func ExecuteConcurrent(img *image.RGBA, pipeline filters.FilterPipeline) {
	var (
		workers     = utils.GetNumberOfWorkers(img.Bounds())
		imageStrips = utils.GetVerticalStrips(img.Bounds(), workers)
		wg sync.WaitGroup
	)

	if pipeline.Preprocess != nil {
		ExecuteConcurrent(img, *pipeline.Preprocess)
	}
	
	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go func(s image.Rectangle) {
			defer wg.Done()
			pipeline.Filter(img, s)
		}(imageStrips[strip])
	}

	wg.Wait()
}
