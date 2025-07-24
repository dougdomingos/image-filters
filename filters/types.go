// Package filters provides a set of image filters implemented using
// pixel-based manipulation techniques.
package filters

import "image"

// Filter is a function that applies a specific effect to the provided image.
// All pixel operations are done in-place, and implementations are thread-safe
// by default.
type Filter func(image *image.RGBA)

// FilterPipeline defines a configurable image processing pipeline, consisting
// of an optional preprocessing stage and the core filter algorithm.
type FilterPipeline struct {

	// Preprocess is an optional step executed before the main filter,
	// enabling recursive chaining of operations. If nil, no preprocessing is
	// performed.
	Preprocess Filter

	Filter Filter
}
