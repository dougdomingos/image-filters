package engines

import (
	"fmt"
	"image"

	"dougdomingos.com/image-filters/filters/types"
)

// ApplyFilterPipeline handles the execution of a filter pipeline, applying
// preprocessing steps recursively (if needed) and selecting the filter
// implementation to be executed.
func ApplyFilterPipeline(img *image.RGBA, pipeline *types.FilterPipeline, isConcurrent bool) error {
	preprocess := pipeline.Preprocess

	if preprocess != nil {
		ApplyFilterPipeline(img, preprocess, isConcurrent)
	}

	filter := getFilterFromPipeline(pipeline, isConcurrent)
	if filter == nil {
		return fmt.Errorf("[ERROR] selected pipeline has no implementation for current execution mode")
	}

	filter(img)
	return nil
}

// getFilterFromPipeline returns the filter function that corresponds to the
// selected execution mode (either serial or concurrent).
func getFilterFromPipeline(pipeline *types.FilterPipeline, isConcurrent bool) types.Filter {
	if isConcurrent {
		return pipeline.ConcurrentFilter
	}
	return pipeline.SerialFilter
}
