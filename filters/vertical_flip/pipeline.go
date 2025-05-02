package vertical_flip

import "dougdomingos.com/image-filters/filters/types"

var VerticalFlipPipeline = types.FilterPipeline{
	Preprocess: nil,
	SerialFilter: serialVerticalFlip,
	ConcurrentFilter: concurrentVerticalFlip,
}