package engines

import (
	"image"

	"dougdomingos.com/image-filters/filters/types"
)

// ApplyFilterPipeline handles the execution of a filter pipeline, applying
// preprocessing steps recursively (if needed) and selecting the filter
// implementation to be executed.
func ApplyFilterPipeline(img *image.RGBA, pipeline *types.FilterPipeline, isConcurrent bool) error {
	if pipeline.Preprocess != nil {
		pipeline.Preprocess(img)
	}

	pipeline.Filter(img)
	return nil
}
