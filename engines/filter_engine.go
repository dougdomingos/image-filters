package engines

import (
	"image"

	"dougdomingos.com/image-filters/filters"
)

// ApplyFilterPipeline handles the execution of a filter pipeline, applying the
// preprocess filter (if present) and then the core filter implementation.
func ApplyFilterPipeline(img *image.RGBA, pipeline *filters.FilterPipeline) {
	if pipeline.Preprocess != nil {
		pipeline.Preprocess(img)
	}

	pipeline.Filter(img)
}
