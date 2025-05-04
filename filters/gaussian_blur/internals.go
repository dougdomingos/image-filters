package gaussian_blur

import (
	"math"
)

const kernelSize = 5
const sigma = 1.5

// GenerateGaussianKernel creates
func generateGaussianKernel(kernelSize int, sigma float64) ([][]float64, int) {
	// if size is even, use the biggest odd number that is smaller than size (fallback == 3)
	if kernelSize%2 == 0 {
		kernelSize = max(3, kernelSize-1)
	}

	sum := 0.0
	centerOffset := kernelSize / 2 // distance from kernel edge to its center

	kernel := make([][]float64, kernelSize)
	for i := range kernel {
		kernel[i] = make([]float64, kernelSize)
	}

	for j := 0; j < kernelSize; j++ {
		for i := range kernelSize {
			x, y := float64(i-centerOffset), float64(j-centerOffset)
			exponent := -((x*x + y*y) / (2 * (sigma * sigma)))
			value := math.Exp(exponent)
			kernel[j][i] = value
			sum += value
		}
	}

	// Normalize the kernel values (i.e., make the sum of all values approx. equal to 1)
	for j := range kernelSize {
		for i := range kernelSize {
			kernel[j][i] /= sum
		}
	}

	return kernel, kernelSize / 2
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
