package engines

import (
	"fmt"
	"image"

	"dougdomingos.com/image-filters/filters/types"
)

func ApplyFilterPipeline(img *image.RGBA, pipeline *types.FilterPipeline, isConcurrent bool) error {
	preprocess := pipeline.Preprocess

	if preprocess != nil {
		ApplyFilterPipeline(img, preprocess, isConcurrent)
	}

	filter := getFilter(pipeline, isConcurrent)
	if filter == nil {
		return fmt.Errorf("pipeline has no implementation for mode \"%s\"", getExecutionMode(isConcurrent))
	}

	filter(img)
	return nil
}

func getFilter(pipeline *types.FilterPipeline, isConcurrent bool) types.Filter {
	var filter types.Filter
	if isConcurrent {
		filter = pipeline.ConcurrentFilter
	} else {
		filter = pipeline.SerialFilter
	}

	return filter
}

func getExecutionMode(isConcurrent bool) string {
	var mode string
	if isConcurrent {
		mode = "concurrent"
	} else {
		mode = "serial"
	}

	return mode
}
