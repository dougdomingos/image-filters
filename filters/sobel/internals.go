package sobel

var (
	// The padding (in pixels) to be added to each side of a image's partition copy.
	copyPadding = 1

	// Sobel's convolution kernel used for computing horizontal gradients.
	gX = [3][3]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	// Sobel's convolution kernel used for computing vertical gradients.
	gY = [3][3]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}
)

// clampColorValue ensures that the provided value stays within the range of an
// RGBA color channel value (i.e., from 0 to 255).
func clampColorValue(val float64) uint8 {
	if val < 0 {
		return 0
	}

	if val > 255 {
		return 255
	}

	return uint8(val)
}
