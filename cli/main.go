package main

import (
	"fmt"
	"path/filepath"

	"dougdomingos.com/image-filters/engines"
	"dougdomingos.com/image-filters/filters"
)

func main() {
	parsedFlags := parseInputFlags()
	
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

	imageFilename := filepath.Base(parsedFlags.ImgPath)
	outputPath, err := SaveImage(imageRGBA, format, parsedFlags.OutputDir, imageFilename)
	if err != nil {
		terminateWithError(err, ImageSavingError)
	}

	fmt.Printf("Processed image stored in \"%s\"\n", outputPath)
}

