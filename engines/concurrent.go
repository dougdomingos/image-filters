package engines

import (
	"image"
	"runtime"
	"sync"

	"dougdomingos.com/image-filters/filters"
)

// ExecuteConcurrent applies the provided filter to the specified image object
// by splitting it into several segments, assigning each one to a goroutine.
func ExecuteConcurrent(img *image.RGBA, pipeline filters.FilterPipeline) {
	var (
		filterToApply = pipeline.Filter
		workers       = getNumberOfWorkers(img.Bounds())
		imageStrips   = pipeline.ConcurrentFilter.PartitionMethod(img.Bounds(), workers)
		wg            sync.WaitGroup
	)

	if pipeline.Preprocess != nil {
		ExecuteConcurrent(img, *pipeline.Preprocess)
	}

	if pipeline.ConcurrentFilter.Builder != nil {
		filterToApply = pipeline.ConcurrentFilter.Builder(img, img.Bounds())
	}

	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go func(s image.Rectangle) {
			defer wg.Done()
			filterToApply(img, s)
		}(imageStrips[strip])
	}

	wg.Wait()
}

// getNumberOfWorkers calculates the optimal number of workers based on the
// number of logical CPUs available and the size of the task (given by the
// bounds). The function considers both the number of available CPUs and the
// width of each worker's segment. It also ensures that, if the width of the
// segments is less than zero, the returned value would be 1.
func getNumberOfWorkers(bounds image.Rectangle) int {
	numLogicCPUs := runtime.NumCPU()
	segmentWidthPerLogicCPU := bounds.Max.X / numLogicCPUs

	return max(min(numLogicCPUs, segmentWidthPerLogicCPU), 1)
}
