package gaussian_blur

import (
	"math"
)

// kernelSize indicates the size of the generated kernel in a NxN format. It is
// required to be a positive odd number greater than or equal to 3.
const kernelSize = 5

// generateGaussianKernel creates a normalized 2D Gaussian kernel of specified
// size and sigma, used to compute the weighted color values of neighbor
// pixels. It returns the kernel matrix and the center offset (kernelSize / 2),
// which is used for aligning the kernel during convolution.
func generateGaussianKernel(size int, sigma float64) ([][]float64, int) {
	kernelSize := ensureValidKernelSize(size)

	kernel := make([][]float64, kernelSize)
	for i := range kernel {
		kernel[i] = make([]float64, kernelSize)
	}

	sum := 0.0
	centerOffset := kernelSize / 2
	for j := range kernelSize {
		for i := range kernelSize {
			x, y := i-centerOffset, j-centerOffset
			value := apply2DGaussianFunction(x, y, sigma)
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

// apply2DGaussianFunction computes the value of the 2D Gaussian distribution at the
// specified (x, y) coordinates with the given standard deviation (sigma).
//
// This function is commonly used in generating Gaussian kernels for image
// processing (e.g., Gaussian blur). The formula is:
//
//   G(x, y) = (1 / sqrt(2πσ²)) * e^(-(x² + y²) / (2σ²))
//
// where σ is the standard deviation, and (x, y) is the distance from the kernel center.
// The output represents the weight or contribution of the pixel at (x, y) to the kernel.
func apply2DGaussianFunction(x, y int, sigma float64) float64 {
	variance := sigma * sigma
	coefficient := 1 / math.Sqrt(2 * math.Pi * variance)
	exponent := -(float64((x * x) + (y * y))) / (2 * variance)
	return coefficient * math.Exp(exponent)
}

// computeKernelSigma applies a heuristic to compute the best standard
// deviation (sigma) to the specified kernel size. The heuristic used is that
// of OpenCV's implementation:
//
//	σ = 0.3 * (((kernelSize - 1) / 2) - 1) + 0.8
//
// σ indicates the standard deviation of the neighbors' weights. Smaller
// values yields more tightly packed kernels (i.e., closer pixels weight more),
// as greater values yields more diffuse kernels (i.e., weights are
// distributed over a larger area).
func computeKernelSigma(size int) float64 {
	kernelSize := float64(ensureValidKernelSize(size))
	return 0.3*((kernelSize-1)*0.5-1) + 0.8
}

// ensureValidKernelSize garantees that the provided kernel size meets the
// constraints of the gaussian kernel.
//
// The kernel size (N) must be a odd number greater than or equal to 3. If N is
// an even number, the function will use the biggest odd number that is smallest
// than N, falling back to N = 3 if the chosen number is still invalid.
func ensureValidKernelSize(size int) int {
	if size%2 != 0 {
		return size
	}

	return max(3, size-1)
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
