package utils

import (
	"flag"
	"fmt"
)

func ParseInputFlags() (string, string, string, int) {
	imgPath := flag.String("img", "", "Path to image file")
	filterName := flag.String("filter", "", "Filter to apply")
	executionMode := flag.String("mode", "serial", "Execution mode: serial or concurrent")
	workers := flag.Int("workers", 1, "Number of threads used to process the image in concurrent mode")
	flag.Parse()

	return *imgPath, *filterName, *executionMode, *workers
}

func ValidateExecutionMode(mode string) error {
	var err error
	if mode != "serial" && mode != "concurrent" {
		err = fmt.Errorf("[ERROR]: Unknown execution model! %s", err)
	}
	
	return err
}
