// Package grayscale implements the grayscale filter.
//
// The grayscale filter converts each pixel of an image into its equivalent
// shade of gray, based on a transform function. This implementation uses the
// Rec. 601 luma transform to compute the grayscale value of each color channel.
// The alpha channel is left unmodified.
package grayscale

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

var GrayscaleFilter = types.NewFilter(nil, Grayscale)

// Grayscale applies the grayscale filter to the entire image using
// multiple goroutines to process different partitions concurrently.
func Grayscale(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetVerticalSegments(bounds, numWorkers)
		wg          sync.WaitGroup
	)

	wg.Add(numWorkers)
	for strip := range imageStrips {
		go grayscaleWorker(img, imageStrips[strip], &wg)
	}

	wg.Wait()
}

// grayscaleWorker processes a partition of the original image by applying the
// grayscale filter to such partition.
func grayscaleWorker(img *image.RGBA, bounds image.Rectangle, wg *sync.WaitGroup) {
	defer wg.Done()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := imgutil.GetRGBA8(img, x, y)
			gray := runLumaTransform(r, g, b)
			imgutil.WriteToPixel(img, x, y, [4]uint8{gray, gray, gray, a})
		}
	}
}
