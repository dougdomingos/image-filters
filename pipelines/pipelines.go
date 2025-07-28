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
	"binarization":     binarization.BinarizationAction,
	"grayscale":        grayscale.GrayscaleAction,
	"horizontal-flip":  horizontal_flip.HorizontalFlipAction,
	"sobel":            sobel.SobelAction,
	"sobel-grayscaled": sobel.SobelGrayscaledAction,
	"vertical-flip":    vertical_flip.VerticalFlipAction,
	"gaussian-blur":    gaussian_blur.GaussianBlurAction,
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
