package imgutil

import (
	"image"
)

// CreatePaddedCopy returns a new RGBA image that is a padded copy of the input
// image. A padding P of black pixels is applied uniformly on all sides of the
// original image, resulting in a new image that is larger by 2 * P in both
// width and height.
//
// Accessing the original image’s pixels within the padded copy requires
// offsetting coordinates by the padding amount: (x + padding, y + padding).
//
// Assumptions:
//   - The original image has bounds starting at (0, 0).
//   - The returned image is safe to share across goroutines for read-only operations.
func CreatePaddedCopy(img image.RGBA, padding int) image.RGBA {
	originalBounds := img.Bounds()

	newMaxX := img.Rect.Max.X + (2 * padding)
	newMaxY := img.Rect.Max.Y + (2 * padding)
	copyImg := image.NewRGBA(image.Rect(0, 0, newMaxX, newMaxY))

	for y := copyImg.Rect.Min.Y; y < copyImg.Rect.Max.Y; y++ {
		copyRow := y * copyImg.Stride

		for x := copyImg.Rect.Min.X; x < copyImg.Rect.Max.X; x++ {
			copyPixOffset := copyRow + (x * 4)

			if !MapsToOriginalPixel(originalBounds, x, y, padding) {
				copy(copyImg.Pix[copyPixOffset:copyPixOffset+4], []uint8{0, 0, 0, 255})
				continue
			}
			
			srcRow := (y - padding) * img.Stride
			srcCol := srcRow + ((x - padding) * 4)
			copy(copyImg.Pix[copyPixOffset:copyPixOffset+4], img.Pix[srcCol:srcCol+4])
		}
	}

	return *copyImg
}

// MapsToOriginalPixel checks whether a given (x, y) coordinate in the padded
// image corresponds to a valid coordinate in the original, unpadded image. It
// assumes that the original image's bounds start at (0, 0) (i.e.,
// bounds.Min.X == 0 and bounds.Min.Y == 0).
func MapsToOriginalPixel(bounds image.Rectangle, x, y, padding int) bool {
	px := x - padding
	py := y - padding

	return (px >= 0 && px < bounds.Dx()) && (py >= 0 && py < bounds.Dy())
}
