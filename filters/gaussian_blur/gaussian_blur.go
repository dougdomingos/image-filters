package gaussian_blur

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
)

// GaussianBlur applies the gaussian blur filter to the entire image
// using multiple goroutines. It computes the global gaussian kernel to be used
// by all workers.
func GaussianBlur(img *image.RGBA) {
	var (
		bounds                       = img.Bounds()
		numWorkers                   = imgutil.GetNumberOfWorkers(bounds)
		imageStrips                  = imgutil.GetVerticalPartitions(bounds, numWorkers)
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

// gaussianBlurWorker process a subregion of the image by creating a padded
// copy of the subregion and using it to compute the weighted color values for
// each pixel within the partition. It ensures that no worker goroutine edits
// certain parts of the original image before the others have finished their
// copy stage.
func gaussianBlurWorker(img, paddedCopy *image.RGBA, bounds image.Rectangle, kernel [][]float64, kernelOffset int, mainWg *sync.WaitGroup) {
	defer mainWg.Done()

	originalBounds := img.Bounds()
	paddedMinX := bounds.Min.X + kernelOffset
	paddedMaxX := bounds.Max.X + kernelOffset
	paddedMinY := bounds.Min.Y + kernelOffset
	paddedMaxY := bounds.Max.Y + kernelOffset

	for y := paddedMinY; y < paddedMaxY; y++ {
		srcRow := (y - kernelOffset) * img.Stride

		for x := paddedMinX; x < paddedMaxX; x++ {
			var sumR, sumG, sumB, sumA, kernelWeightSum float64
			srcCol := srcRow + (x-kernelOffset)*4

			for ky := -kernelOffset; ky <= kernelOffset; ky++ {
				for kx := -kernelOffset; kx <= kernelOffset; kx++ {
					deltaX, deltaY := x+kx, y+ky

					// disconsider padding pixels from blurring calculations
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

			updatedColor := []uint8{clamp256(sumR), clamp256(sumG), clamp256(sumB), clamp256(sumA)}
			copy(img.Pix[srcCol:srcCol+4], updatedColor)
		}
	}
}
