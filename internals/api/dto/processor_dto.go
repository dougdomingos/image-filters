package dto

import (
	"fmt"
	"image"

	"dougdomingos.com/image-filters/pipelines"
)

// ProcessorRequestDTO represents the expected parameters for a parsed
// processor request.
type ProcessorRequestDTO struct {

	// Img is the decoded image.
	Img image.Image

	// ImgFormat is the inferred format of the decoded image.
	ImgFormat string

	// ImgFilename is the name of the image file sent through the request.
	ImgFilename string

	// Pipeline represents the list of filters to be applied to the image.
	Pipeline pipelines.Pipeline
}

// ProcessorResponseDTO represents the JSON for processor HTTP responses.
type ProcessorResponseDTO struct {

	// ProcessedImageURL is the remote path to access the processed image.
	ProcessedImageURL string `json:"processed_image_url"`
}

// BuildProcessorResponse constructs a ProcessorResponseDTO with the remote URL
// for the processed image, stored in the server.
func BuildProcessorResponse(imgFilename string) ProcessorResponseDTO {
	imageURL := fmt.Sprintf("/images/%s", imgFilename)

	return ProcessorResponseDTO{
		ProcessedImageURL: imageURL,
	}
}
