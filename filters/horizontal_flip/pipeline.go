// Package horizontal_flip implements the horizontal flip filter.
//
// The horizontal flip filter reverses the order of pixels in each row,
// effectively mirroring the image along its vertical axis.
package horizontal_flip

import "dougdomingos.com/image-filters/filters/types"

// HorizontalFlipPipeline defines the horizontal flip filter pipeline. As this
// filter does not require any preprocessing pipeline, the preprocess step is
// not declared.
var HorizontalFlipPipeline = types.FilterPipeline{
	Preprocess:       nil,
	SerialFilter:     serialHorizontalFlip,
	ConcurrentFilter: concurrentHorizontalFlip,
}
