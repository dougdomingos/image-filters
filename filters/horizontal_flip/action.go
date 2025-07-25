// Package horizontal_flip implements the horizontal flip filter.
//
// The horizontal flip filter reverses the order of pixels in each row,
// effectively mirroring the image along its vertical axis.
package horizontal_flip

import (
	"dougdomingos.com/image-filters/filters"
)

var HorizontalFlipAction = filters.Action{
	Preprocess: nil,
	Filter:     HorizontalFlip,
}
