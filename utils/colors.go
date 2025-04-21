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