package filters

import "fmt"

// AvailableFilters maps a string identifier to its corresponding filter
// pipeline.
var AvaliableFilters = map[string]FilterPipeline{
	"binarization":     BinarizationPipeline,
	"grayscale":        GrayscalePipeline,
	"horizontal-flip":  HorizontalFlipPipeline,
	"sobel":            SobelPipeline,
	"sobel-grayscaled": SobelGrayscaledPipeline,
	"vertical-flip":    VerticalFlipPipeline,
	// add more filters here...
}

// GetFilterPipeline retrieves a filter pipeline by its name from the
// AvailableFilters map. It returns the pipeline if found, or an error if the
// specified name is not defined.
func GetFilterPipeline(filterName string) (FilterPipeline, error) {
	pipeline, exists := AvaliableFilters[filterName]
	if !exists {
		return FilterPipeline{}, fmt.Errorf("[ERROR]: Specified filter does not exist")
	}

	return pipeline, nil
}
