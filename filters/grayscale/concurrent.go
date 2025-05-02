package grayscale

import (
	"image"
	"sync"

	"dougdomingos.com/image-filters/partitions"
	"dougdomingos.com/image-filters/utils"
)

// concurrentGrayscale converts each pixel within a specified segment of the image into
// its corresponding shade of gray.
//
// This implementation uses the Rec. 601 luma transform to compute the
// grayscale value. The alpha channel is left unmodified.
func concurrentGrayscale(img *image.RGBA) {
	var (
		bounds = img.Bounds()
		numWorkers = 8
		imageStrips = partitions.GetVerticalPartitions(bounds, numWorkers)
		wg sync.WaitGroup
	)

	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go grayscaleWorker(img, imageStrips[strip], &wg)
	}

	wg.Wait()
}

func grayscaleWorker(img *image.RGBA, bounds image.Rectangle, wg *sync.WaitGroup) {
	defer wg.Done()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := (y - img.Rect.Min.Y) * img.Stride
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := rowStart + (x-img.Rect.Min.X)*4

			r, g, b, a := utils.GetRGBA8(img, x, y)
			gray := RunLumaTransform(r, g, b)

			copy(img.Pix[offset:offset+4], []uint8{gray, gray, gray, a})
		}
	}
}