package gaussian_blur

import (
	"math"
)

// kernelSize indicates the size of the generated kernel in a NxN format. It is
// required to be a positive odd number greater than or equal to 3. 
const kernelSize = 5

// sigma indicates the standard deviation of the neighbors' weights. Smaller
// values yields more tightly packed kernels (i.e., closer pixels weight more),
// as greater values yields more diffuse kernels (i.e., weights are
// distributed over a larger area).
const sigma = 1.5

// generateGaussianKernel creates a normalized 2D Gaussian kernel of specified
// size and sigma, used to compute the weighted color values of neighbor
// pixels. It returns the kernel matrix and the center offset (kernelSize / 2),
// which is used for aligning the kernel during convolution.
//
// The kernel size (N) must be a odd number greater than or equal to 3. If N is
// an even number, the function will use the biggest odd number that is smallest
// than N, falling back to N = 3 if the chosen number is still invalid.
func generateGaussianKernel(kernelSize int, sigma float64) ([][]float64, int) {
	if kernelSize%2 == 0 {
		kernelSize = max(3, kernelSize-1)
	}

	kernel := make([][]float64, kernelSize)
	for i := range kernel {
		kernel[i] = make([]float64, kernelSize)
	}

	sum := 0.0
	centerOffset := kernelSize / 2
	for j := range kernelSize {
		for i := range kernelSize {
			x, y := float64(i-centerOffset), float64(j-centerOffset)
			exponent := -((x*x + y*y) / (2 * (sigma * sigma)))
			value := math.Exp(exponent)
			kernel[j][i] = value
			sum += value
		}
	}

	// normalize the kernel values (i.e., make the sum of all values approx. equal to 1)
	for j := range kernelSize {
		for i := range kernelSize {
			kernel[j][i] /= sum
		}
	}

	return kernel, centerOffset
}

// clamp256 ensures that the provided value stays within the range of an
// RGBA color channel value (i.e., from 0 to 255).
func clamp256(val float64) uint8 {
	if val < 0 {
		return 0
	}

	if val > 255 {
		return 255
	}

	return uint8(val)
}
