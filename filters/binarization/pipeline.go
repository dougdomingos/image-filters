// Package binarization implements the binarization filter.
//
// The binarization filter converts an image to black and white based on a
// brightness threshold, which can either be fixed or calculated based on the
// image's pixels.
//
// This implementation applies Otsu's Method of thresholding to determine the
// value that best separates the foreground and background components of the
// image.
package binarization

import (
	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/types"
)

// BinarizationPipeline defines the binarization filter pipeline. Since
// binarization works best on grayscaled image, it uses the GrayscalePipeline
// as a preprocessing step.
var BinarizationPipeline = types.FilterPipeline{
	Preprocess:       &grayscale.GrayscalePipeline,
	SerialFilter:     serialBinarization,
	ConcurrentFilter: concurrentBinarization,
}
