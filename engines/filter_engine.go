package engines

import (
	"image"

	"dougdomingos.com/image-filters/pipelines"
)

func ProcessRecipe(img *image.RGBA, recipe *pipelines.Recipe) {
	for step := recipe.Head; step != nil; step = recipe.NextStep() {
		if step.Action.Preprocess != nil {
			step.Action.Preprocess(img)
		}

		step.Action.Filter(img)
	}
}
