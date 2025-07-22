package services

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
)

func convertImageToRGBA(img image.Image) *image.RGBA {
	if rgba, ok := img.(*image.RGBA); ok {
		return rgba
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return rgba
}

func encodeResponseImage(w http.ResponseWriter, img *image.RGBA, format string) {
	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", format))

	switch format {
	case "png":
		png.Encode(w, img)
	case "jpeg":
		jpeg.Encode(w, img, nil)
	}
}
