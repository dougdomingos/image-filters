// Package grayscale implements the grayscale filter.
//
// The grayscale filter converts each pixel of an image into its equivalent
// shade of gray, based on a transform function. This implementation uses the
// Rec. 601 luma transform to compute the grayscale value of each color channel.
// The alpha channel is left unmodified.
package grayscale

import "dougdomingos.com/image-filters/filters/types"

var GrayscaleFilter = types.NewFilter(nil, Grayscale)