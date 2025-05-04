package gaussian_blur

import "dougdomingos.com/image-filters/filters/types"

var GaussianBlurPipeline = types.FilterPipeline{
	Preprocess:       nil,
	SerialFilter:     serialGaussianBlur,
	ConcurrentFilter: concurrentGaussianBlur,
}
