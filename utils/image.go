package utils

import "image"

// GetRGBA8 extracts the channel values from a specific pixel within an image
// and returns them as 8-bit unsigned integers, ranging from 0 to 255.
func GetRGBA8(img *image.RGBA, x, y int) (r, g, b, a uint8) {
	red, green, blue, alpha := img.At(x, y).RGBA()

	red8 := uint8(red >> 8)
	green8 := uint8(green >> 8)
	blue8 := uint8(blue >> 8)
	alpha8 := uint8(alpha >> 8)

	return red8, green8, blue8, alpha8
}

// SwapPixels swaps two pixels' positions within a RGBA image. It's based on
// Go's RGBA implementation, in which pixels channels are disposed as a
// unidimensional array of unsigned 8-bit integers.
func SwapPixels(img *image.RGBA, offsetX, offsetY int) {
	for i := range 4 {
		img.Pix[offsetX+i], img.Pix[offsetY+i] = img.Pix[offsetY+i], img.Pix[offsetX+i]
	}
}

// CopyImagePartition creates a copy of a region of the provided image.
func CopyImagePartition(img *image.RGBA, bounds image.Rectangle) image.RGBA {
	copiedPartition := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			srcOffset := img.PixOffset(x, y)
			dstOffset := copiedPartition.PixOffset(x, y)

			copy(copiedPartition.Pix[dstOffset:dstOffset+4], img.Pix[srcOffset:srcOffset+4])
		}
	}

	return *copiedPartition
}

func CopyImagePartitionWithPadding(img *image.RGBA, bounds image.Rectangle, padding int) image.RGBA {
	paddedBounds := GetPaddedBounds(bounds, padding)
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

func GetPaddedBounds(bounds image.Rectangle, padding int) image.Rectangle {
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
