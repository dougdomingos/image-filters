package grayscale

// RunLumaTransform computes the grayscale value of a pixel based on the Rec. 601
// luma transform:
//
//	gray = (0.299 * R) + (0.587 * G) + (0.114 * B)
//
// The Rec. 601 standard accounts for human visual perception, assigning higher
// weight to the green channel,followed by red and blue, based on the eye's
// sensitivity to each color.
func RunLumaTransform(r, g, b uint8) uint8 {
	weightedRed := 0.299 * float64(r)
	weightedGreen := 0.587 * float64(g)
	weightedBlue := 0.114 * float64(b)

	return uint8(weightedRed + weightedGreen + weightedBlue)
}
