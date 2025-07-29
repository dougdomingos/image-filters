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
var AvaliableFilters = map[string]types.Filter{
	"binarization":     binarization.BinarizationFilter,
	"grayscale":        grayscale.GrayscaleFilter,
	"horizontal-flip":  horizontal_flip.HorizontalFlipFilter,
	"sobel":            sobel.SobelFilter,
	"sobel-grayscaled": sobel.SobelGrayscaledFilter,
	"vertical-flip":    vertical_flip.VerticalFlipFilter,
	"gaussian-blur":    gaussian_blur.GaussianBlurFilter,
	// add more filters here...
}

// GetFilter retrieves a filter by its name from the AvailableFilters map.
// It returns the filter, if present, or an error if the specified filter is
// not defined.
func GetFilter(filterID string) (types.Filter, error) {
	pipeline, exists := AvaliableFilters[filterID]
	if !exists {
		return types.Filter{}, fmt.Errorf("[ERROR]: Specified filter does not exist")
	}

	return pipeline, nil
}

// GetAvaliableFilterIDs returns a slice containing the identifiers of all
// filters declared in the AvaliableFilters map.
func GetAvaliableFilterIDs() []string {
	filterIDs := make([]string, 0, len(AvaliableFilters))

	for filterKey := range AvaliableFilters {
		filterIDs = append(filterIDs, filterKey)
	}

	return filterIDs
}
