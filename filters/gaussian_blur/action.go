// Pacakge gaussian_blur implements the gaussian blur filter.
//
// The Gaussian blur filter applies a smoothing effect to an image by averaging
// neighboring pixels with a Gaussian function, which reduces image noise and
// detail. The kernel used in color calculations is automatically generated
// based on a predefined kernel size and standard deviation of the neighbors
// (represented by sigma, or "σ").
package gaussian_blur

import "dougdomingos.com/image-filters/filters"

var GaussianBlurAction = filters.Action{
	Preprocess: nil,
	Filter:     GaussianBlur,
}
