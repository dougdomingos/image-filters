package dto

import (
	"image"

	"dougdomingos.com/image-filters/filters/types"
)

// ProcessorRequestDTO represents the expected parameters for a parsed
// processor request.
type ProcessorRequestDTO struct {

	// Img is the decoded image.
	Img          image.Image

	// ImgFormat is the inferred format of the decoded image.
	ImgFormat    string

	// Pipeline is the requested filter pipeline that matches the "filter"
	// parameter from the request.
	Pipeline     types.FilterPipeline

	// IsConcurrent signals the engine to use the thread-safe pipeline
	// implementation.
	IsConcurrent bool
}
