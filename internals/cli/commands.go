package cli

import (
	"fmt"
	"sort"
	"time"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/internals/utils"
	"dougdomingos.com/image-filters/pipelines"
)

// ListAvaliablePipelines displays the list of all avaliable pipelines.
func ListAvaliablePipelines() {
	pipelineIDs := make([]string, 0, len(pipelines.AvaliableFilters))

	for filterKey := range pipelines.AvaliableFilters {
		pipelineIDs = append(pipelineIDs, filterKey)
	}

	sort.Strings(pipelineIDs)

	fmt.Println("Avaliable pipelines:")
	for _, pipeline := range pipelineIDs {
		fmt.Printf("\t => %s\n", pipeline)
	}
}

func ApplyRecipeToImage(imgPath, outputDir string, filters []string) {
	imageRGBA, format, err := utils.LoadImage(imgPath)
	if err != nil {
		terminateWithError(err, ImageLoadingError)
	}

	err = utils.CreateOutputDir(outputDir)
	if err != nil {
		terminateWithError(err, OutputDirError)
	}

	recipe, err := pipelines.BuildRecipe(filters)
	if err != nil {
		terminateWithError(err, FilterNotFoundError)
	}

	engines.ProcessRecipe(imageRGBA, &recipe)

	outputFile := utils.GetProcessedImageFilename(imgPath, time.Now().String())
	outputPath, err := utils.SaveImage(imageRGBA, format, outputDir, outputFile)
	if err != nil {
		terminateWithError(err, ImageSavingError)
	}

	fmt.Printf("Processed image stored in \"%s\"\n", outputPath)
}
