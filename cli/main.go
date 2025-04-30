package main

import (
	"fmt"
	"image"
	"path/filepath"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/utils"
)

func main() {
	parsedFlags := parseInputFlags()

	err := utils.CreateOutputDir(parsedFlags.OutputDir)
	if err != nil {
		terminateWithError(err, OutputDirError)
	}

	err = validateExecutionMode(parsedFlags.ExecutionMode)
	if err != nil {
		terminateWithError(err, InvalidExecutionModeError)
	}

	pipeline, err := filters.GetFilterPipeline(parsedFlags.FilterName)
	if err != nil {
		terminateWithError(err, FilterNotFoundError)
	}

	imageRGBA, format, err := utils.LoadImage(parsedFlags.ImgPath)
	if err != nil {
		terminateWithError(err, ImageLoadingError)
	}

	applyFilter(imageRGBA, pipeline, parsedFlags.ExecutionMode)

	imageFilename := filepath.Base(parsedFlags.ImgPath)
	outputPath, err := utils.SaveImage(imageRGBA, format, parsedFlags.OutputDir, imageFilename)
	if err != nil {
		terminateWithError(err, ImageSavingError)
	}

	fmt.Printf("Processed image stored in \"%s\"\n", outputPath)
}

// applyFilter executes the given filter on the image using the specified mode.
func applyFilter(imageRGBA *image.RGBA, pipeline filters.FilterPipeline, mode string) {
	switch mode {
	case "serial":
		engines.ExecuteSerial(imageRGBA, pipeline)
	case "concurrent":
		engines.ExecuteConcurrent(imageRGBA, pipeline)
	}
}
