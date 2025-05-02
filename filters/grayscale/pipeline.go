// Package grayscale implements the grayscale filter.
//
// The grayscale filter converts each pixel of an image into its equivalent
// shade of gray, based on a transform function. This implementation uses the
// Rec. 601 luma transform to compute the grayscale value of each color channel.
// The alpha channel is left unmodified.
package grayscale

import "dougdomingos.com/image-filters/filters/types"

// GrayscalePipeline defines the grayscale filter pipeline. Grayscale does not
// require any preprocessing pipeline, and so, no preprocess step is declared.
var GrayscalePipeline = types.FilterPipeline{
	Preprocess:       nil,
	SerialFilter:     serialGrayscale,
	ConcurrentFilter: concurrentGrayscale,
}
