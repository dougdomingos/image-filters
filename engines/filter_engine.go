package engines

import (
	"fmt"
	"image"

	"dougdomingos.com/image-filters/pipelines"
)

// ProcessPipeline sequentially applies all filters defined in the given
// pipeline to the provided image.
//
// If the provided pipeline is nil, empty or contains a PipelineStep with a nil
// Filter, processing stops and an error is returned. Otherwise, returns nil.
func ProcessPipeline(img *image.RGBA, pipeline *pipelines.Pipeline) error {
	if pipeline == nil || pipeline.IsEmpty() {
		return fmt.Errorf("[ERROR] unable to process empty pipeline")
	}

	for !pipeline.IsEmpty() {
		currentStep := pipeline.NextStep()
		if currentStep.Filter == nil {
			return fmt.Errorf("[ERROR] pipeline step has no filter")
		}

		currentStep.Filter.ApplyFilter(img)
	}

	return nil
}
