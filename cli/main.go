package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	inputImage := flag.String("img", "", "Path to image file")
	filterName := flag.String("filter", "", "Filter to apply")
	executionMode := flag.String("mode", "serial", "Execution mode: serial or concurrent")
	workers := flag.Int("workers", 1, "Number of threads used to process the image in concurrent mode")
	flag.Parse()

	if *executionMode == "serial" {
		// this shall work, eventually...
		fmt.Printf("Running on %s mode...\n", *executionMode)
	} else if *executionMode == "concurrent" {
		// and so will this.
		fmt.Printf("Running on %s mode with %d worker(s)...\n", *executionMode, *workers)
	} else {
		fmt.Printf("[ERROR] Unknown execution mode: %s\n", *executionMode)
		os.Exit(1)
	}

	fmt.Printf("Image file: \"%s\", filter \"%s\"\n", *inputImage, *filterName)
}
