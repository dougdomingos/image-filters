// Package sobel implements an image filter that detects edges in images using
// the Sobel operator, a common edge detection algorithm in image processing.
//
// The Sobel operator works by applying two 3x3 convolution kernels (for
// horizontal and vertical edges, respectively) to compute the gradient
// magnitude of each pixel. The resulting gradient magnitude highlights regions
// with significant brightness changes, thus highlighting edges in the image.
package sobel

import (
	"image"
	"math"
	"sync"

	"dougdomingos.com/image-filters/filters/grayscale"
	"dougdomingos.com/image-filters/filters/imgutil"
	"dougdomingos.com/image-filters/filters/types"
)

// SobelFilter applies the sobel transformation to an image. No pre-processing
// is performed.
var SobelFilter = types.NewFilter(nil, Sobel)

// SobelGrayscaledFilter applies a grayscale pre-process, followed by the sobel
// transformation. As the Sobel operator relies on pixel brightness, it works
// best on grayscaled images.
var SobelGrayscaledFilter = types.NewFilter(grayscale.Grayscale, Sobel)

// Sobel applies the sobel transformation to the entire image. The image is
// divided into vertical segments, each delegated to a worker goroutine.
// 
// Each goroutine also receives a global, padded copy of the image, which
// ensures that kernel positions are always valid and preserving original data
// until all computations are finished.
func Sobel(img *image.RGBA) {
	var (
		bounds      = img.Bounds()
		numWorkers  = imgutil.GetNumberOfWorkers(bounds)
		imageStrips = imgutil.GetVerticalSegments(bounds, numWorkers)
		paddedCopy  = imgutil.CreatePaddedCopy(*img, copyPadding)
		mainWg      sync.WaitGroup
	)

	mainWg.Add(numWorkers)
	for strip := range imageStrips {
		go sobelWorker(img, &paddedCopy, imageStrips[strip], &mainWg)
	}

	mainWg.Wait()
}

// sobelWorker computes Sobel gradients for a segment of the image. It uses a
// padded copy of the original to safely access neighboring pixels. The
// gradients of each pixel are computed based on the brightness of its
// neighbors, where greater variations lead to more intense edges.
func sobelWorker(srcImg, paddedCopy *image.RGBA, bounds image.Rectangle, mainWg *sync.WaitGroup) {
	defer mainWg.Done()

	paddedMinX := bounds.Min.X + copyPadding
	paddedMaxX := bounds.Max.X + copyPadding
	paddedMinY := bounds.Min.Y + copyPadding
	paddedMaxY := bounds.Max.Y + copyPadding

	for y := paddedMinY; y < paddedMaxY; y++ {
		for x := paddedMinX; x < paddedMaxX; x++ {
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

			updatedPixel := [4]uint8{gradR, gradG, gradB, a8}
			imgutil.WriteToPixel(srcImg, x-copyPadding, y-copyPadding, updatedPixel)
		}
	}
}
