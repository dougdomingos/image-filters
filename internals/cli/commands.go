package cli

import (
	"fmt"
	"time"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/internals/utils"
	"dougdomingos.com/image-filters/pipelines"
)

// ListAvaliablePipelines displays the list of all avaliable pipelines.
func ListAvaliablePipelines() {
	filterIDs := filters.GetAvaliableFilterIDs()

	fmt.Println("Avaliable pipelines:")
	for _, pipeline := range filterIDs {
		fmt.Printf("\t => %s\n", pipeline)
	}
}

func ApplyPipelineToImage(imgPath, outputDir string, filters []string) {
	imageRGBA, format, err := utils.LoadImage(imgPath)
	if err != nil {
		terminateWithError(err, ImageLoadingError)
	}

	err = utils.CreateOutputDir(outputDir)
	if err != nil {
		terminateWithError(err, OutputDirError)
	}

	pipeline, err := pipelines.NewPipeline(filters)
	if err != nil {
		terminateWithError(err, FilterNotFoundError)
	}

	if err := engines.ProcessPipeline(imageRGBA, &pipeline); err != nil {
		terminateWithError(err, MalformedPipelineError)
	}

	outputFile := utils.GetProcessedImageFilename(imgPath, time.Now().String())
	outputPath, err := utils.SaveImage(imageRGBA, format, outputDir, outputFile)
	if err != nil {
		terminateWithError(err, ImageSavingError)
	}

	fmt.Printf("Processed image stored in \"%s\"\n", outputPath)
}
