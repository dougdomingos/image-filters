package gaussian_blur

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/utils"
)

func concurrentGaussianBlur(img *image.RGBA) {
	var (
		bounds                       = img.Bounds()
		numWorkers                   = imgutil.GetNumberOfWorkers(bounds)
		imageStrips                  = imgutil.GetVerticalPartitions(bounds, numWorkers)
		gaussianKernel, kernelOffset = generateGaussianKernel(kernelSize, sigma)
		mainWg                       sync.WaitGroup
		copyWg                       sync.WaitGroup
	)

	mainWg.Add(numWorkers)
	copyWg.Add(numWorkers)
	for strip := range imageStrips {
		go gaussianBlurWorker(img, imageStrips[strip], gaussianKernel, kernelOffset, &mainWg, &copyWg)
	}

	mainWg.Wait()
}

func gaussianBlurWorker(img *image.RGBA, bounds image.Rectangle, kernel [][]float64, kernelOffset int, mainWg, copyWg *sync.WaitGroup) {
	defer mainWg.Done()
	copyImg := func() image.RGBA {
		defer copyWg.Done()
		return imgutil.CopyPaddedImagePartition(img, bounds, kernelOffset)
	}()

	paddedMinX, paddedMaxX := copyImg.Rect.Min.X+kernelOffset, copyImg.Rect.Max.X-kernelOffset
	paddedMinY, paddedMaxY := copyImg.Rect.Min.Y+kernelOffset, copyImg.Rect.Max.Y-kernelOffset
	for y := paddedMinY; y < paddedMaxY; y++ {
		srcRowStart := (y - kernelOffset) * img.Stride
		for x := paddedMinX; x < paddedMaxX; x++ {
			if x == paddedMaxX-kernelOffset {
				copyWg.Wait()
			}

			var sumR, sumG, sumB, sumA, kernelWeightSum float64

			for ky := -kernelOffset; ky <= kernelOffset; ky++ {
				for kx := -kernelOffset; kx <= kernelOffset; kx++ {
					deltaX, deltaY := x+kx, y+ky

					if !imgutil.IsPositionWithinOriginalImage(img, deltaX, deltaY, kernelOffset) {
						continue
					}

					r, g, b, a := utils.GetRGBA8(&copyImg, deltaX, deltaY)
					kernelWeight := kernel[ky+kernelOffset][kx+kernelOffset]

					sumR += float64(r) * kernelWeight
					sumG += float64(g) * kernelWeight
					sumB += float64(b) * kernelWeight
					sumA += float64(a) * kernelWeight
					kernelWeightSum += kernelWeight
				}
			}

			if kernelWeightSum > 0 {
				sumR /= kernelWeightSum
				sumG /= kernelWeightSum
				sumB /= kernelWeightSum
				sumA /= kernelWeightSum
			}

			updatedColor := []uint8{clamp256(sumR), clamp256(sumG), clamp256(sumB), clamp256(sumA)}
			copyOffset := srcRowStart + (x-kernelOffset)*4
			copy(img.Pix[copyOffset:copyOffset+4], updatedColor)
		}
	}
}
