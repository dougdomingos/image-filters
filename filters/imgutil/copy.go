package imgutil

import "image"

func CopyPaddedImagePartition(img *image.RGBA, bounds image.Rectangle, padding int) image.RGBA {
	paddedBounds := getPaddedBounds(bounds, padding)
	paddedImg := image.NewRGBA(paddedBounds)

	for y := paddedBounds.Min.Y; y < paddedBounds.Max.Y; y++ {
		copyRowStart := y * paddedImg.Stride
		for x := paddedBounds.Min.X; x < paddedBounds.Max.X; x++ {
			copyOffset := copyRowStart + (x-paddedImg.Rect.Min.X)*4

			if !IsPositionWithinOriginalImage(img, x, y, padding) {
				copy(paddedImg.Pix[copyOffset:copyOffset+4], []uint8{0, 0, 0, 255})
			} else {
				srcRowStart := (y - padding) * img.Stride
				srcOffset := srcRowStart + (x-padding-img.Rect.Min.X)*4
				copy(paddedImg.Pix[copyOffset:copyOffset+4], img.Pix[srcOffset:srcOffset+4])
			}
		}
	}

	return *paddedImg
}

func IsPositionWithinOriginalImage(img *image.RGBA, x, y, padding int) bool {
	offsetX, offsetY := x-padding, y-padding

	isXWithinBounds := offsetX >= img.Rect.Min.X && offsetX < img.Rect.Max.X
	isYWithinBounds := offsetY >= img.Rect.Min.Y && offsetY < img.Rect.Max.Y

	return isXWithinBounds && isYWithinBounds
}

func getPaddedBounds(bounds image.Rectangle, padding int) image.Rectangle {
	minX, maxX := bounds.Min.X-padding, bounds.Max.X+padding
	minY, maxY := bounds.Min.Y-padding, bounds.Max.Y+padding

	if bounds.Min.X == 0 {
		minX, maxX = 0, bounds.Max.X+(2*padding)
	}

	if bounds.Min.Y == 0 {
		minY, maxY = 0, bounds.Max.Y+(2*padding)
	}

	return image.Rect(minX, minY, maxX, maxY)
}
