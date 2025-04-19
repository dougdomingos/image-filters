package main

import (
	"fmt"
	"flag"
)

func main() {
	inputImage := flag.String("img", "", "Path to image file")
	filterName := flag.String("filter", "", "Filter to apply")
	flag.Parse()

	fmt.Printf("Image file: \"%s\", filter \"%s\"\n", *inputImage, *filterName)
}