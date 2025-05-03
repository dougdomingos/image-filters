package binarization

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/utils"
)

// concurrentBinarization applies the binarization filter to the entire image
// using multiple goroutines. It first computes Otsu's global threshold,
// then partitions the image and processes each partition concurrently.
func concurrentBinarization(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetVerticalPartitions(bounds, numWorkers)
		threshold   = otsuThreshold(img, bounds)
		wg          sync.WaitGroup
	)

	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go binarizationWorker(img, imageStrips[strip], threshold, &wg)
	}

	wg.Wait()
}

// binarizationWorker processes a subregion of the image by applying the
// binarization filter based on a shared global threshold. It updates each
// pixel in the subregion to either black or white, depending on it's
// intensity.
func binarizationWorker(img *image.RGBA, bounds image.Rectangle, threshold uint8, wg *sync.WaitGroup) {
	defer wg.Done()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := (y - img.Rect.Min.Y) * img.Stride
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := rowStart + (x-img.Rect.Min.X)*4
			intensity, _, _, _ := utils.GetRGBA8(img, x, y)

			if intensity > threshold {
				copy(img.Pix[offset:offset+4], whitePixel[:])
			} else {
				copy(img.Pix[offset:offset+4], blackPixel[:])
			}

		}
	}
}
