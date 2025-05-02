// Package filters provides a set of image filters implemented using
// pixel-based manipulation techniques.
package types

import "image"

// Filter represents a function that applies a specific image processing
// operation. The operation is performed in-place, modifying the input image
// directly to reduce memory overhead by avoiding unnecessary copies.
type Filter func(image *image.RGBA)

// FilterPipeline defines a configurable image processing pipeline, consisting
// of an optional preprocessing stage and two implementations of the main
// filter algorithm: serial and concurrent.
//
// The pipeline supports recursive composition, allowing complex operations
// to be built by chaining simpler FilterPipeline instances.
type FilterPipeline struct {

	// Preprocess is an optional step executed before the main filter,
	// enabling recursive chaining of operations. If nil, no preprocessing is
	// performed.
	Preprocess *FilterPipeline

	// SerialFilter is the main implementation of the filter algorithm,
	// that process the entire image in a single pass.
	SerialFilter Filter

	// ConcurrentFilter is a variation of the main filter that supports parallel
	// processing. It splits the image into different partitions and processes
	// them independently.
	ConcurrentFilter Filter
}
