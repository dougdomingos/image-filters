package sobel

import (
	"image"
	"math"
	"sync"

	"dougdomingos.com/image-filters/filters/imgutil"
)

// Sobel applies the Sobel filter to the entire image using multiple
// goroutines. It makes a full copy of the image, which is shared amongst the
// routines, which then modify their respective partitions.
func Sobel(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetVerticalPartitions(bounds, numWorkers)
		paddedCopy  = imgutil.CreatePaddedCopy(*img, copyPadding)
		mainWg      sync.WaitGroup
	)

	mainWg.Add(numWorkers)
	for strip := range imageStrips {
		go sobelWorker(img, &paddedCopy, imageStrips[strip], &mainWg)
	}

	mainWg.Wait()
}

// sobelWorker processes a subregion of the image by applying the sobel filter
// based on a global copy of the original image, computing the kernel values of
// each color channel of each pixel in the subregion.
func sobelWorker(srcImg, paddedCopy *image.RGBA, bounds image.Rectangle, mainWg *sync.WaitGroup) {
	defer mainWg.Done()

	paddedMinX := bounds.Min.X + copyPadding
	paddedMaxX := bounds.Max.X + copyPadding
	paddedMinY := bounds.Min.Y + copyPadding
	paddedMaxY := bounds.Max.Y + copyPadding

	for y := paddedMinY; y < paddedMaxY; y++ {
		srcRow := (y - copyPadding) * srcImg.Stride

		for x := paddedMinX; x < paddedMaxX; x++ {
			srcCol := srcRow + (x-copyPadding)*4

			var r8, g8, b8, a8 uint8
			var gxR, gxG, gxB, gyR, gyG, gyB int

			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					deltaX := x + kx
					deltaY := y + ky

					r8, g8, b8, a8 = imgutil.GetRGBA8(paddedCopy, deltaX, deltaY)
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

			copy(srcImg.Pix[srcCol:srcCol+4], []uint8{gradR, gradG, gradB, a8})
		}
	}
}
