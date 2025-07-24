package services

import (
	"image"
	"image/draw"
)

// convertImageToRGBA takes a image passed through a multipart/form-data request
// and converts it to a editable RGBA format, if necessary.
func convertImageToRGBA(img image.Image) *image.RGBA {
	if rgba, ok := img.(*image.RGBA); ok {
		return rgba
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return rgba
}
