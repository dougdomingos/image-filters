package dto

import (
	"image"

	"dougdomingos.com/image-filters/filters/types"
)

type ProcessorRequestDTO struct {
	Img          image.Image
	ImgFormat    string
	Pipeline     types.FilterPipeline
	IsConcurrent bool
}
