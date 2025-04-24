package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/utils"
)

func main() {
	imagePath, outputDir, filterName, executionMode := utils.ParseInputFlags()

	err := utils.CreateOutputDir(outputDir)
	if err != nil {
		terminateWithError(err, OutputDirError)
	}

	err = utils.ValidateExecutionMode(executionMode)
	if err != nil {
		terminateWithError(err, InvalidExecutionModeError)
	}

	filter, err := filters.GetFilter(filterName)
	if err != nil {
		terminateWithError(err, FilterNotFoundError)
	}

	imageRGBA, format, err := utils.LoadImage(imagePath)
	if err != nil {
		terminateWithError(err, ImageLoadingError)
	}

	applyFilter(imageRGBA, filter, executionMode)
	imageFilename := filepath.Base(imagePath)

	outputPath, err := utils.SaveImage(imageRGBA, format, outputDir, imageFilename)
	if err != nil {
		terminateWithError(err, ImageSavingError)
	}

	fmt.Printf("Processed image stored in \"%s\"\n", outputPath)
}

// applyFilter executes the given filter on the image using the specified mode.
func applyFilter(imageRGBA *image.RGBA, filter filters.Filter, mode string) {
	switch mode {
	case "serial":
		engines.ExecuteSerial(imageRGBA, filter)
	case "concurrent":
		engines.ExecuteConcurrent(imageRGBA, filter)
	}
}

// terminateWithError prints the error and exits the program with the given code.
func terminateWithError(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}
