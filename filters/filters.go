// Package filters provides a set of image filters implemented using
// pixel-based manipulation techniques.
package filters

import (
	"fmt"

	"dougdomingos.com/image-filters/filters/binarization"
	"dougdomingos.com/image-filters/filters/gaussian_blur"
	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/horizontal_flip"
	"dougdomingos.com/image-filters/filters/sobel"
	"dougdomingos.com/image-filters/filters/types"
	"dougdomingos.com/image-filters/filters/vertical_flip"
)

// AvailableFilters maps a string identifier to its corresponding filter
// action.
var AvaliableFilters = map[string]types.Action{
	"binarization":     binarization.BinarizationAction,
	"grayscale":        grayscale.GrayscaleAction,
	"horizontal-flip":  horizontal_flip.HorizontalFlipAction,
	"sobel":            sobel.SobelAction,
	"sobel-grayscaled": sobel.SobelGrayscaledAction,
	"vertical-flip":    vertical_flip.VerticalFlipAction,
	"gaussian-blur":    gaussian_blur.GaussianBlurAction,
	// add more actions here...
}

// GetFilterAction retrieves a filter pipeline by its name from the
// AvailableFilters map. It returns the pipeline if found, or an error if the
// specified name is not defined.
func GetFilterAction(filterID string) (types.Action, error) {
	pipeline, exists := AvaliableFilters[filterID]
	if !exists {
		return types.Action{}, fmt.Errorf("[ERROR]: Specified filter does not exist")
	}

	return pipeline, nil
}
