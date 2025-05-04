package main

import (
	"flag"

	"dougdomingos.com/image-filters/filters"
)

// ParsedFlags declares a structure used for retrieving the flag values
// provided by the user.
type ParsedFlags struct {
	ImgPath       string
	OutputDir     string
	FilterName    string
	IsConcurrent  bool
	ListPipelines bool
}

// parseInputFlags declares the expected command-line flags and parses their
// values into a ParsedFlags struct for later usage.
func parseInputFlags() ParsedFlags {
	imgPath := flag.String("img", "", "Path to image file")
	outputDir := flag.String("outDir", "./output", "Directory where the processed image shall be stored")
	filterName := flag.String("filter", "", "Name of the filter pipeline to be applied")
	isConcurrent := flag.Bool("concurrent", false, "Specify if filter pipeline should be use parallel execution")
	listPipelines := flag.Bool("list", false, "List all avaliable filter pipelines")
	flag.Parse()

	return ParsedFlags{
		ImgPath:       *imgPath,
		OutputDir:     *outputDir,
		FilterName:    *filterName,
		IsConcurrent:  *isConcurrent,
		ListPipelines: *listPipelines,
	}
}

// getAvaliablePipelines returns the string identifiers of all avaliable filter
// pipelines.
func getAvaliablePipelines() []string {
	keys := make([]string, 0, len(filters.AvaliableFilters))

	for filterKey := range filters.AvaliableFilters {
		keys = append(keys, filterKey)
	}

	return keys
}
