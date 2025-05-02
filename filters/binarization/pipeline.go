package binarization

import (
	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/types"
)

var BinarizationPipeline = types.FilterPipeline{
	Preprocess:       &grayscale.GrayscalePipeline,
	SerialFilter:     serialBinarization,
	ConcurrentFilter: concurrentBinarization,
}
