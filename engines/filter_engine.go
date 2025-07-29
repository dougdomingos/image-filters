package engines

import (
	"image"

	"dougdomingos.com/image-filters/pipelines"
)

func ProcessPipeline(img *image.RGBA, pipeline *pipelines.Pipeline) {
	if pipeline.IsEmpty() {
		return
	}

	currentStep := pipeline.Peek()
	for !pipeline.IsEmpty() {
		currentStep.Filter.ApplyFilter(img)
	}
}
