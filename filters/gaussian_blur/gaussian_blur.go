// Pacakge gaussian_blur implements an image smoothing filter based on a
// gaussian convolution kernel. It reduces noise and detail by blending each
// pixel with its neighbors through a gaussian kernel, where closer pixels are
// more significant. The strength of the blur effect depends on the kernel size
// and standard deviation (σ).
package gaussian_blur

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

// GaussianBlurFilter applies the gaussian blur transformation to an image. No
// pre-processing is performed.
var GaussianBlurFilter = types.NewFilter(nil, GaussianBlur)

// GaussianBlur applies the gaussian blur transformation to the entire image.
// The image is divided into vertical segments, each delegated to a worker
// goroutine.
// 
// Each goroutine also receives a global, padded copy of the image, which
// ensures that kernel positions are always valid and preserving original data
// until all computations are finished.
func GaussianBlur(img *image.RGBA) {
	var (
		bounds                       = img.Bounds()
		numWorkers                   = imgutil.GetNumberOfWorkers(bounds)
		imageStrips                  = imgutil.GetVerticalSegments(bounds, numWorkers)
		sigma                        = computeKernelSigma(kernelSize)
		gaussianKernel, kernelOffset = generateGaussianKernel(kernelSize, sigma)
		paddedCopy                   = imgutil.CreatePaddedCopy(*img, kernelOffset)
		mainWg                       sync.WaitGroup
	)

	mainWg.Add(numWorkers)
	for strip := range imageStrips {
		go gaussianBlurWorker(img, &paddedCopy, imageStrips[strip], gaussianKernel, kernelOffset, &mainWg)
	}

	mainWg.Wait()
}

// gaussianBlurWorker applies the Gaussian kernel to a segment of the image. It
// uses a padded copy of the original to safely access neighboring pixels.
func gaussianBlurWorker(img, paddedCopy *image.RGBA, bounds image.Rectangle, kernel [][]float64, kernelOffset int, mainWg *sync.WaitGroup) {
	defer mainWg.Done()

	originalBounds := img.Bounds()
	paddedMinX := bounds.Min.X + kernelOffset
	paddedMaxX := bounds.Max.X + kernelOffset
	paddedMinY := bounds.Min.Y + kernelOffset
	paddedMaxY := bounds.Max.Y + kernelOffset

	for y := paddedMinY; y < paddedMaxY; y++ {
		for x := paddedMinX; x < paddedMaxX; x++ {
			var sumR, sumG, sumB, sumA, kernelWeightSum float64

			for ky := -kernelOffset; ky <= kernelOffset; ky++ {
				for kx := -kernelOffset; kx <= kernelOffset; kx++ {
					deltaX, deltaY := x+kx, y+ky

					// skip padding pixels from gradient calculations
					if !imgutil.MapsToOriginalPixel(originalBounds, deltaX, deltaY, kernelOffset) {
						continue
					}

					r, g, b, a := imgutil.GetRGBA8(paddedCopy, deltaX, deltaY)
					kernelWeight := kernel[ky+kernelOffset][kx+kernelOffset]

					sumR += float64(r) * kernelWeight
					sumG += float64(g) * kernelWeight
					sumB += float64(b) * kernelWeight
					sumA += float64(a) * kernelWeight
					kernelWeightSum += kernelWeight
				}
			}

			// normalize the weighted sums to mitigate loss on edge píxels
			if kernelWeightSum > 0 {
				sumR /= kernelWeightSum
				sumG /= kernelWeightSum
				sumB /= kernelWeightSum
				sumA /= kernelWeightSum
			}

			updatedColor := [4]uint8{clamp256(sumR), clamp256(sumG), clamp256(sumB), clamp256(sumA)}
			imgutil.WriteToPixel(img, x-kernelOffset, y-kernelOffset, updatedColor)
		}
	}
}
