package horizontal_flip

import "dougdomingos.com/image-filters/filters/types"

var HorizontalFlipPipeline = types.FilterPipeline{
	Preprocess:       nil,
	SerialFilter:     serialHorizontalFlip,
	ConcurrentFilter: concurrentHorizontalFlip,
}
