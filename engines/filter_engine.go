package engines

import (
	"image"

	"dougdomingos.com/image-filters/pipelines"
)

func ProcessRecipe(img *image.RGBA, recipe *pipelines.Recipe) {
	currentStep := recipe.NextStep()
	for currentStep != nil {
		if currentStep.Action.Preprocess != nil {
			currentStep.Action.Preprocess(img)
		}

		currentStep.Action.Filter(img)
		currentStep = recipe.NextStep()
	}
}
