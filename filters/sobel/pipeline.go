package sobel

import (
	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/types"
)

var SobelPipeline = types.FilterPipeline{
	Preprocess:       nil,
	SerialFilter:     serialSobel,
	ConcurrentFilter: concurrentSobel,
}

var SobelGrayscaledPipeline = types.FilterPipeline{
	Preprocess:       &grayscale.GrayscalePipeline,
	SerialFilter:     serialSobel,
	ConcurrentFilter: concurrentSobel,
}
