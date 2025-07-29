package main

import (
	"dougdomingos.com/image-filters/internals/cli"
)

func main() {
	flags := cli.ParseInputFlags()

	if flags.ListPipelines {
		cli.ListAvaliablePipelines()
	} else {
		cli.ApplyPipelineToImage(flags.ImgPath, flags.OutputDir, flags.Filters)
	}
}
