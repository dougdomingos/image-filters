package services

import (
	"encoding/json"
	"image"
	"image/draw"
	"net/http"
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

// sendJSONResponse encodes the provided data as a JSON object to be returned
// as response for a request.
func sendJSONResponse(w http.ResponseWriter, responseData any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
