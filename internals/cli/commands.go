package cli

import (
	"fmt"
	"sort"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
)

// ListAvaliablePipelines displays the list of all avaliable pipelines.
func ListAvaliablePipelines() {
	pipelineIDs := make([]string, 0, len(filters.AvaliableFilters))

	for filterKey := range filters.AvaliableFilters {
		pipelineIDs = append(pipelineIDs, filterKey)
	}

	sort.Strings(pipelineIDs)

	fmt.Println("Avaliable pipelines:")
	for _, pipeline := range pipelineIDs {
		fmt.Printf("\t => %s\n", pipeline)
	}
}

// ApplyPipelineToImage loads an image from disk, applies the selected filter
// pipeline to it and store the result image into the specified output
// directory. It also allows the selection of the execution mode.
func ApplyPipelineToImage(imgPath, outputDir, pipelineID string, isConcurrent bool) {
	imageRGBA, format, err := loadImage(imgPath)
	if err != nil {
		terminateWithError(err, ImageLoadingError)
	}

	err = createOutputDir(outputDir)
	if err != nil {
		terminateWithError(err, OutputDirError)
	}

	pipeline, err := filters.GetFilterPipeline(pipelineID)
	if err != nil {
		terminateWithError(err, FilterNotFoundError)
	}

	err = engines.ApplyFilterPipeline(imageRGBA, &pipeline, isConcurrent)
	if err != nil {
		terminateWithError(err, FilterNotImplementedError)
	}

	outputFile := getProcessedImageFilename(imgPath, pipelineID)
	outputPath, err := saveImage(imageRGBA, format, outputDir, outputFile)
	if err != nil {
		terminateWithError(err, ImageSavingError)
	}

	fmt.Printf("Processed image stored in \"%s\"\n", outputPath)
}
