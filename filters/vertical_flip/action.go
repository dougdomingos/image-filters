// Package vertical_flip implements the vertical flip filter.
//
// The vertical flip filter reverses the order of pixels in each column,
// effectively mirroring the image along its horizontal axis.
package vertical_flip

import (
	"dougdomingos.com/image-filters/filters/types"
)

var VerticalFlipFilter = types.NewFilter(nil, VerticalFlip)
