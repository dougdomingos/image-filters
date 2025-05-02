package filters

import (
	"fmt"

	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/types"
)

// AvailableFilters maps a string identifier to its corresponding filter
// pipeline.
var AvaliableFilters = map[string]types.FilterPipeline{
	"binarization":     BinarizationPipeline,
	"grayscale":        grayscale.GrayscalePipeline,
	"horizontal-flip":  HorizontalFlipPipeline,
	"sobel":            SobelPipeline,
	"sobel-grayscaled": SobelGrayscaledPipeline,
	"vertical-flip":    VerticalFlipPipeline,
	// add more filters here...
}

// GetFilterPipeline retrieves a filter pipeline by its name from the
// AvailableFilters map. It returns the pipeline if found, or an error if the
// specified name is not defined.
func GetFilterPipeline(filterName string) (types.FilterPipeline, error) {
	pipeline, exists := AvaliableFilters[filterName]
	if !exists {
		return types.FilterPipeline{}, fmt.Errorf("[ERROR]: Specified filter does not exist")
	}

	return pipeline, nil
}
