// Package filters provides a set of image filters implemented using
// pixel-based manipulation techniques.
package filters

import (
	"fmt"
	"image"
)

// Filter is a function that applies a certain processing technique within a
// defined segment of a image.
//
// Filters manipulate an image in-place, which avoid the need to create copies
// of the original image when dealing with multithread processing and merge
// then once the operations are finished.
//
// The `image` is the full in-memory image, and the `bounds` is the region of
// the image that the filter should operate on. The filter directly alters
// `img`.
type Filter func(image *image.RGBA, bounds image.Rectangle)

// AvailableFilters maps a string identifier to its corresponding image filter
// function.
var AvaliableFilters = map[string]Filter{
	"grayscale": Grayscale,
	"binarization": Binarization,
	// add more filters here...
}

// GetFilter retrieves a filter function by its name from the AvailableFilters
// map. It returns the filter if found, or an error if the filter name is not
// defined.
func GetFilter(filterName string) (Filter, error) {
	filter, exists := AvaliableFilters[filterName]
	if !exists {
		return nil, fmt.Errorf("[ERROR]: Specified filter does not exist")
	}

	return filter, nil
} 