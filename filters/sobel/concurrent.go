package sobel

import (
	"image"
	"math"
	"sync"

	"dougdomingos.com/image-filters/filters/partitioning"
	"dougdomingos.com/image-filters/utils"
)

// concurrentSobel applies the sobel filter to the entire image using multiple
// goroutines. It makes a full copy of the image, which is shared amongst the
// routines, which then modify their respective partitions.
func concurrentSobel(img *image.RGBA) {
	var (
		padding     = 1
		bounds      = img.Bounds()
		numWorkers  = partitioning.GetNumberOfWorkers(bounds)
		imageStrips = partitioning.GetVerticalPartitions(bounds, numWorkers)
		wg          sync.WaitGroup
	)

	copyImg := utils.CopyImagePartitionWithPadding(img, bounds, padding)
	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go sobelWorker(img, imageStrips[strip], &copyImg, padding, &wg)
	}

	wg.Wait()
}

// sobelWorker processes a subregion of the image by applying the sobel filter
// based on a global copy of the original image, which is used to compute the
// gradients of each color channel.
func sobelWorker(img *image.RGBA, bounds image.Rectangle, copyImg *image.RGBA, padding int, wg *sync.WaitGroup) {
	defer wg.Done()

	paddedMinX, paddedMaxX := copyImg.Rect.Min.X+padding, copyImg.Rect.Max.X-padding
	paddedMinY, paddedMaxY := copyImg.Rect.Min.Y+padding, copyImg.Rect.Max.Y-padding

	for y := paddedMinY; y < paddedMaxY; y++ {
		srcRowStart := (y - padding) * img.Stride
		for x := paddedMinX; x < paddedMaxX; x++ {
			offset := srcRowStart + (x-padding)*4

			var r8, g8, b8, a8 uint8
			var gxR, gxG, gxB, gyR, gyG, gyB int

			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					deltaX := x + kx
					deltaY := y + ky

					r8, g8, b8, a8 = utils.GetRGBA8(copyImg, deltaX, deltaY)
					r := int(r8)
					g := int(g8)
					b := int(b8)

					valKernelX := gX[ky+1][kx+1]
					valKernelY := gY[ky+1][kx+1]

					gxR += int(valKernelX * r)
					gxG += int(valKernelX * g)
					gxB += int(valKernelX * b)

					gyR += int(valKernelY * r)
					gyG += int(valKernelY * g)
					gyB += int(valKernelY * b)
				}
			}

			gradR := clampColorValue(math.Sqrt(float64(gxR*gxR + gyR*gyR)))
			gradG := clampColorValue(math.Sqrt(float64(gxG*gxG + gyG*gyG)))
			gradB := clampColorValue(math.Sqrt(float64(gxB*gxB + gyB*gyB)))

			copy(img.Pix[offset:offset+4], []uint8{gradR, gradG, gradB, a8})
		}
	}
}
