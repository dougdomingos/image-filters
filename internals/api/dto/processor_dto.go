package dto

import (
	"fmt"
	"image"

	"dougdomingos.com/image-filters/filters/types"
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

	// FilterName is the ID of the requested filter pipeline, used to
	// identify the filter that was applied to a image persisted on disk.
	FilterName string

	// Pipeline is the requested filter pipeline that matches the "filter"
	// parameter from the request.
	Pipeline types.FilterPipeline

	// IsConcurrent signals the engine to use the thread-safe pipeline
	// implementation.
	IsConcurrent bool
}

// ProcessorResponseDTO represents the JSON for processor HTTP responses.
type ProcessorResponseDTO struct {

	// ProcessedImageURL is the remote path to access the processed image.
	ProcessedImageURL string `json:"processedImageURL"`
}

// BuildProcessorResponse constructs a ProcessorResponseDTO with the remote URL
// for the processed image, stored in the server.
func BuildProcessorResponse(imgFilename string) ProcessorResponseDTO {
	imageURL := fmt.Sprintf("/images/%s", imgFilename)

	return ProcessorResponseDTO{
		ProcessedImageURL: imageURL,
	}
}
