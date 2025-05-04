package main

import (
	"fmt"
	"os"
	"sort"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
)

func main() {
	parsedFlags := parseInputFlags()

	if parsedFlags.ListPipelines {
		pipelineIDs := getAvaliablePipelines()
		sort.Strings(pipelineIDs)

		fmt.Println("Avaliable pipelines:")
		for _, pipeline := range pipelineIDs {
			fmt.Printf("\t => %s\n", pipeline)
		}

		os.Exit(0)
	}

	imageRGBA, format, err := LoadImage(parsedFlags.ImgPath)
	if err != nil {
		terminateWithError(err, ImageLoadingError)
	}

	err = CreateOutputDir(parsedFlags.OutputDir)
	if err != nil {
		terminateWithError(err, OutputDirError)
	}

	pipeline, err := filters.GetFilterPipeline(parsedFlags.FilterName)
	if err != nil {
		terminateWithError(err, FilterNotFoundError)
	}

	err = engines.ApplyFilterPipeline(imageRGBA, &pipeline, parsedFlags.IsConcurrent)
	if err != nil {
		terminateWithError(err, FilterNotImplementedError)
	}

	outputFile := GetProcessedImageFilename(parsedFlags.ImgPath, parsedFlags.FilterName)
	outputPath, err := SaveImage(imageRGBA, format, parsedFlags.OutputDir, outputFile)
	if err != nil {
		terminateWithError(err, ImageSavingError)
	}

	fmt.Printf("Processed image stored in \"%s\"\n", outputPath)
}
