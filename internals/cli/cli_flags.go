package cli

import "flag"

// ParsedFlags declares a structure used for retrieving the flag values
// provided by the user.
type ParsedFlags struct {
	ImgPath       string
	OutputDir     string
	FilterName    string
	IsConcurrent  bool
	ListPipelines bool
}

// ParseInputFlags declares the expected command-line flags and parses their
// values into a ParsedFlags struct for later usage.
func ParseInputFlags() ParsedFlags {
	imgPath := flag.String("img", "", "Path to image file")
	outputDir := flag.String("outDir", "./output", "Directory where the processed image shall be stored")
	filterName := flag.String("filter", "", "Name of the filter pipeline to be applied")
	isConcurrent := flag.Bool("concurrent", false, "Specify if filter pipeline should be use parallel execution")
	listPipelines := flag.Bool("listPipelines", false, "Specify if filter pipeline should be use parallel execution")

	flag.Parse()

	return ParsedFlags{
		ImgPath:      *imgPath,
		OutputDir:    *outputDir,
		FilterName:   *filterName,
		IsConcurrent: *isConcurrent,
		ListPipelines: *listPipelines,
	}
}
