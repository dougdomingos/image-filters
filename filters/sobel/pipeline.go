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

// SobelPipeline defines the sobel filter pipeline. As Sobel's Operator does
// not require any preprocessing to detect edges within a image, no
// preprocessing step is declared.
var SobelPipeline = types.FilterPipeline{
	Preprocess:       nil,
	SerialFilter:     serialSobel,
	ConcurrentFilter: concurrentSobel,
}

// SobelGrayscaledPipeline defines the sobel filter pipeline with a grayscale
// preprocess step. Since Sobel's Operator relies on color variation, it may
// be hindered by color noise (e.g., detecting high variations of one color
// channel, but not the other). As such, its results can be improved by using
// grayscaled images, which would only account variations on brightness and
// intensity.
var SobelGrayscaledPipeline = types.FilterPipeline{
	Preprocess:       &grayscale.GrayscalePipeline,
	SerialFilter:     serialSobel,
	ConcurrentFilter: concurrentSobel,
}
