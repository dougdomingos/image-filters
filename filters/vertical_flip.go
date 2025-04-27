package filters

import (
	"image"

	"dougdomingos.com/image-filters/partitions"
	"dougdomingos.com/image-filters/utils"
)

// VerticalFlipPipeline defines the vertical flip filter pipeline.
//
// This pipeline includes the following:
//   - Preprocess: none
//   - BuildConcurrent: none; works on serial and concurrent modes.
//     Requires vertical segmentation for concurrent mode.
var VerticalFlipPipeline = BuildFilterPipeline(VerticalFlip, nil, nil, partitions.GetVerticalPartitions)

// VerticalFlip reverses the order of each column of pixels within the
// specified segment of the image.
func VerticalFlip(img *image.RGBA, bounds image.Rectangle) {
	middle := (bounds.Max.Y - bounds.Min.Y) / 2

	for deltaY := range middle {
		rowStartTop := (bounds.Min.Y + deltaY) * img.Stride
		rowStartBottom := (bounds.Max.Y - 1 - deltaY) * img.Stride

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			topOffset := rowStartTop + (x * 4)
			bottomOffset := rowStartBottom + (x * 4)

			utils.SwapPixels(img, topOffset, bottomOffset)
		}
	}
}
