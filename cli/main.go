package main

import (
	"fmt"
	"os"

	"dougdomingos.com/image-filters/utils"
)

func main() {
	inputImage, filterName, executionMode := utils.ParseInputFlags()

	fmt.Printf("Running on %s mode...\n", executionMode)
	if executionMode == "serial" {
		// this shall work, eventually...
	} else if executionMode == "concurrent" {
		// and so will this.
	} else {
		fmt.Printf("[ERROR] Unknown execution mode: %s\n", executionMode)
		os.Exit(1)
	}

	fmt.Printf("Image file: \"%s\", filter \"%s\"\n", inputImage, filterName)
}
