// Package grayscale implements an image filter that converts colored images
// to grayscale, based on perceived brightness of the colors of each pixel. 
//
// This implementation uses the Rec. 601 luma transform, which weights color
// channels according to human brightness perception, producing more natural
// results. The alpha channel remains unchanged.
package grayscale

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

// GrayscaleFilter applies the grayscale transformation to a image. No
// pre-processing is performed.
var GrayscaleFilter = types.NewFilter(nil, Grayscale)

// Grayscale applies the grayscale transformation to the entire image. The
// image is divided into vertical segments, each delegated to a worker
// goroutine.
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

// grayscaleWorker converts all pixels in the given image segment to grayscale
// by computing their luminance and updating them in-place.
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
