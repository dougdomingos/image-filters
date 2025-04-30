package main

import (
	"flag"
	"fmt"
)

// ParsedFlags declares a structure used for retrieving the flag values
// provided by the user.
type ParsedFlags struct {
	ImgPath       string
	OutputDir     string
	FilterName    string
	ExecutionMode string
}

// parseInputFlags declares the expected command-line flags and parses their
// values into a ParsedFlags struct for later usage.
func parseInputFlags() ParsedFlags {
	imgPath := flag.String("img", "", "Path to image file")
	outputDir := flag.String("outputDir", "./output", "Directory where the processed image shall be stored")
	filterName := flag.String("filter", "", "Filter to apply")
	executionMode := flag.String("mode", "serial", "Execution mode: serial or concurrent")
	flag.Parse()

	return ParsedFlags{
		ImgPath:       *imgPath,
		OutputDir:     *outputDir,
		FilterName:    *filterName,
		ExecutionMode: *executionMode,
	}
}

// validateExecutionMode ensures that the provided execution mode is either
// "serial" or "concurrent"
func validateExecutionMode(mode string) error {
	var err error
	if mode != "serial" && mode != "concurrent" {
		err = fmt.Errorf("[ERROR]: Unknown execution model! %s", err)
	}

	return err
}
