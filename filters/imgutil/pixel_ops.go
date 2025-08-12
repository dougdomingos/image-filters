package imgutil

import "image"

// GetRGBA8 extracts the channel values from a specific pixel within an image
// and returns them as 8-bit unsigned integers, ranging from 0 to 255.
func GetRGBA8(img *image.RGBA, x, y int) (r, g, b, a uint8) {
	offset := img.PixOffset(x, y)

	red := img.Pix[offset]
	green := img.Pix[offset+1]
	blue := img.Pix[offset+2]
	alpha := img.Pix[offset+3]

	return red, green, blue, alpha
}

// WriteToPixel sets the pixel at (x, y) in the image to the specified RGBA
// values. The pixel array must contain red, green, blue, and alpha components
// in that order. The caller must ensure (x, y) is within the bounds of the
// provided image.
func WriteToPixel(img *image.RGBA, x, y int, pixel [4]uint8) {
	offset := img.PixOffset(x, y)
	copy(img.Pix[offset:offset+4], pixel[:])
}

// SwapPixels exchanges the RGBA values of the pixels at (x1, y1) and (x2, y2)
// in the specified image. The caller must ensure the coordinates are within
// the bounds of the provided image.
func SwapPixels(img *image.RGBA, x1, y1, x2, y2 int) {
	offset1 := img.PixOffset(x1, y1)
	offset2 := img.PixOffset(x2, y2)

	for i := range 4 {
		img.Pix[offset1+i], img.Pix[offset2+i] = img.Pix[offset2+i], img.Pix[offset1+i]
	}
}
