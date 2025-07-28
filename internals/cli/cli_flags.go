package cli

import (
	"flag"
	"strings"
)

// ParsedFlags declares a structure used for retrieving the flag values
// provided by the user.
type ParsedFlags struct {
	ImgPath       string
	OutputDir     string
	Filters       []string
	ListPipelines bool
}

// ParseInputFlags declares the expected command-line flags and parses their
// values into a ParsedFlags struct for later usage.
func ParseInputFlags() ParsedFlags {
	imgPath := flag.String("img", "", "Path to image file")
	outputDir := flag.String("outDir", "./output", "Directory where the processed image shall be stored")
	filters := flag.String("filters", "", "Name of the filter pipeline to be applied")
	listPipelines := flag.Bool("listPipelines", false, "Specify if filter pipeline should be use parallel execution")

	flag.Parse()

	return ParsedFlags{
		ImgPath:       *imgPath,
		OutputDir:     *outputDir,
		Filters:       parseMultiValueFlag(filters),
		ListPipelines: *listPipelines,
	}
}

func parseMultiValueFlag(rawValue *string) []string {
	if *rawValue == "" {
		return nil
	}

	// Split and clean the input
	filters := strings.Split(*rawValue, ",")
	for i := range filters {
		filters[i] = strings.TrimSpace(filters[i])
	}

	return filters
}
