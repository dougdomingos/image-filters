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

// SwapPixels swaps two pixels' positions within a RGBA image. It's based on
// Go's RGBA implementation, in which pixels channels are disposed as a
// unidimensional array of unsigned 8-bit integers.
func SwapPixels(img *image.RGBA, offsetX, offsetY int) {
	for i := range 4 {
		img.Pix[offsetX+i], img.Pix[offsetY+i] = img.Pix[offsetY+i], img.Pix[offsetX+i]
	}
}
