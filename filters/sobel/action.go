// Package sobel implements the sobel filter.
//
// The sobel filter applies Sobel's Operator for edge detection to measure the
// existence and intensity of edges within the image. Edges are detected by
// measuring the color variation between a pixel and its neighbors. Sobel's
// Operator relies on two convolution kernels to detect edges in the vertical
// and horizontal axis.
package sobel

import (
	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/types"
)

var SobelFilter = types.NewFilter(nil, Sobel)
var SobelGrayscaledFilter = types.NewFilter(grayscale.Grayscale, Sobel)
