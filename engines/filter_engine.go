package engines

import (
	"image"

	"dougdomingos.com/image-filters/pipelines"
)

// ApplyFilterPipeline handles the execution of a filter pipeline, applying the
// preprocess filter (if present) and then the core filter implementation.
func ApplyFilterPipeline(img *image.RGBA, recipe *pipelines.Recipe) {
	currentStep := recipe.NextStep()
	for currentStep != nil {
		if currentStep.Action.Preprocess != nil {
			currentStep.Action.Preprocess(img)
		}

		currentStep.Action.Filter(img)
		currentStep = recipe.NextStep()
	}
}
