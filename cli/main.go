package main

import (
	"fmt"
	"os"

	"dougdomingos.com/image-filters/utils"
)

func main() {
	inputImage, filterName, executionMode, workers := utils.ParseInputFlags()

	if executionMode == "serial" {
		// this shall work, eventually...
		fmt.Printf("Running on %s mode...\n", executionMode)
	} else if executionMode == "concurrent" {
		// and so will this.
		fmt.Printf("Running on %s mode with %d worker(s)...\n", executionMode, workers)
	} else {
		fmt.Printf("[ERROR] Unknown execution mode: %s\n", executionMode)
		os.Exit(1)
	}

	fmt.Printf("Image file: \"%s\", filter \"%s\"\n", inputImage, filterName)
}
