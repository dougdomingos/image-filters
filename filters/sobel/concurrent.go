package sobel

import (
	"image"
	"math"
	"sync"

	"dougdomingos.com/image-filters/partitions"
	"dougdomingos.com/image-filters/utils"
)

func concurrentSobel(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = 8
		imageStrips = partitions.GetVerticalPartitions(bounds, numWorkers)
		wg          sync.WaitGroup
	)

	copyImg := utils.CopyImagePartition(img, bounds)
	wg.Add(len(imageStrips))
	for strip := range imageStrips {
		go SobelWorker(img, imageStrips[strip], &copyImg, &wg)
	}

	wg.Wait()
}

func SobelWorker(img *image.RGBA, bounds image.Rectangle, copyImg *image.RGBA, wg *sync.WaitGroup) {
	defer wg.Done()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		rowStart := (y - img.Rect.Min.Y) * img.Stride
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			offset := rowStart + (x-img.Rect.Min.X)*4

			var r8, g8, b8, a8 uint8
			var gxR, gxG, gxB, gyR, gyG, gyB int

			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					deltaX := x + kx
					deltaY := y + ky

					// If the current neighbor is outside of the image, jump to next iteration
					if !isPositionWithinImage(img, deltaX, deltaY) {
						continue
					}

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
