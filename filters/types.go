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
// The image is the full in-memory image, and the bounds is the region of the
// image that the filter should operate on.
type Filter func(image *image.RGBA, bounds image.Rectangle)

// FilterPipeline defines the entire processing flow of a filter, composed of
// an optional preprocessing step, the core filter and an optional closure
// function to adapt filters for concurrent workflows.
//
// A pipeline can recursively chain preprocessing steps, making it flexible for
// building complex operations by composing a series of simpler ones.
type FilterPipeline struct {

	// Preprocess defines an optional preprocessing pipeline to be applied
	// once to the entire image before the main filter runs. It can itself be
	// another FilterPipeline, allowing for recursive chaining of operations.
	// If nil, no preprocessing operations are performed.
	Preprocess *FilterPipeline

	// Filter is the core filter applied in both serial (by default) and
	// concurrent execution modes. For concurrent execution, only the image
	// boundaries should be modified.
	Filter Filter

	// BuildConcurrent, if non-nil, is called once with the full image bounds
	// to return a custom Filter suitable for concurrent execution (e.g.
	// injecting a global threshold). If nil, the Filter is assumed to be safe
	// for concurrent use as-is.
	BuildConcurrent func(img *image.RGBA, bounds image.Rectangle) Filter
}

// AvailableFilters maps a string identifier to its corresponding filter
// pipeline.
var AvaliableFilters = map[string]FilterPipeline{
	"grayscale":    GrayscalePipeline,
	"binarization": BinarizationPipeline,
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
