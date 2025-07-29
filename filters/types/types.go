package types

import "image"

// Transformation defines a single image operation that modifies the provided
// instance in-place. Implementations are expected to be thread-safe.
type Transformation func(img *image.RGBA)

// Filter defines an image transformation composed of optional pre-processing
// and required processing steps.
type Filter struct {

	// preprocess applies an optional pre-processing transformation to the
	// image prior to the main processing step. If nil, no pre-processing is
	// applied.
	preprocess Transformation

	// apply performs the main transformation on the image. This function must
	// not be nil.
	apply Transformation
}

// ApplyFilter applies the filter to the provided image. If a pre-processing
// function is defined, it is applied first, followed by the main transformation.
//
// The operation is performed in-place and modifies the original image.
func (filter *Filter) ApplyFilter(img *image.RGBA) {
	if filter.preprocess != nil {
		filter.preprocess(img)
	}

	filter.apply(img)
}

// NewFilter constructs a new Filter from the given pre-processing and main
// processing functions. The main processing function must be non-nil.
//
// If 'main' is nil, the returned Filter will be a zero value (invalid).
func NewFilter(pre, main Transformation) Filter {
	if main == nil {
		panic("BuildFilter: 'main' processing function must not be nil")
	}

	return Filter{
		preprocess: pre,
		apply:      main,
	}
}
