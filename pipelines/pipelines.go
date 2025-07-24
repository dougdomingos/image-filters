package pipelines

import (
	"fmt"

	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/filters/binarization"
	"dougdomingos.com/image-filters/filters/gaussian_blur"
	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/horizontal_flip"
	"dougdomingos.com/image-filters/filters/sobel"
	"dougdomingos.com/image-filters/filters/vertical_flip"
)

// AvailableFilters maps a string identifier to its corresponding filter
// pipeline.
var AvaliableFilters = map[string]filters.Action{
	"binarization":     binarization.BinarizationPipeline,
	"grayscale":        grayscale.GrayscalePipeline,
	"horizontal-flip":  horizontal_flip.HorizontalFlipPipeline,
	"sobel":            sobel.SobelPipeline,
	"sobel-grayscaled": sobel.SobelGrayscaledPipeline,
	"vertical-flip":    vertical_flip.VerticalFlipPipeline,
	"gaussian-blur":    gaussian_blur.GaussianBlurPipeline,
	// add more filters here...
}

// GetFilterPipeline retrieves a filter pipeline by its name from the
// AvailableFilters map. It returns the pipeline if found, or an error if the
// specified name is not defined.
func GetFilterPipeline(filterName string) (filters.Action, error) {
	pipeline, exists := AvaliableFilters[filterName]
	if !exists {
		return filters.Action{}, fmt.Errorf("[ERROR]: Specified filter does not exist")
	}

	return pipeline, nil
}
