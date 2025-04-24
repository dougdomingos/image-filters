package utils

import (
	"flag"
	"fmt"
)

func ParseInputFlags() (string, string, string, string) {
	imgPath := flag.String("img", "", "Path to image file")
	outputDir := flag.String("outputDir", "./output", "Directory where the processed image shall be stored")
	filterName := flag.String("filter", "", "Filter to apply")
	executionMode := flag.String("mode", "serial", "Execution mode: serial or concurrent")
	flag.Parse()

	return *imgPath, *outputDir, *filterName, *executionMode
}

func ValidateExecutionMode(mode string) error {
	var err error
	if mode != "serial" && mode != "concurrent" {
		err = fmt.Errorf("[ERROR]: Unknown execution model! %s", err)
	}
	
	return err
}
