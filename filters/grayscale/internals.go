package grayscale

// runLumaTransform computes the brightness value of a pixel based on the
// Rec. 601 luma transform:
//
//	brightness = (0.299 * R) + (0.587 * G) + (0.114 * B)
//
// The Rec. 601 standard accounts for human perception of brightness, assigning
// a higher weight to green, followed by red and blue.
func runLumaTransform(r, g, b uint8) uint8 {
	weightedRed := 0.299 * float64(r)
	weightedGreen := 0.587 * float64(g)
	weightedBlue := 0.114 * float64(b)

	return uint8(weightedRed + weightedGreen + weightedBlue)
}
