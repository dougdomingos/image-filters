package engines

import (
	"image"
	"dougdomingos.com/image-filters/filters"
)

// ExecuteSerial applies the provided filter to the specified image object
// as a whole, without any sort of partition. As such, all processing is
// done with a single thread.
func ExecuteSerial(image *image.RGBA, filter filters.Filter) {
	filter(image, image.Bounds())
}