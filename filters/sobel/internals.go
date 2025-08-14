package sobel

var (
	// copyPadding defines the padding size added around the copy of the image.
	// This padding ensures safety when computing edge pixels.
	copyPadding = 1

	// gX is Sobel's horizontal edge detection kernel.
	gX = [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	// gY is Sobel's vertical edge detection kernel.
	gY = [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}
)

// clampColorValue clips a float64 value to a valid 8-bit color channel, ranged
// from 0 to 255.
func clampColorValue(val float64) uint8 {
	if val < 0 {
		return 0
	}

	if val > 255 {
		return 255
	}

	return uint8(val)
}
