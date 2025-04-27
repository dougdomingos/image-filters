// Package filters provides a set of image filters implemented using
// pixel-based manipulation techniques.
package filters

import (
	"image"

	"dougdomingos.com/image-filters/partitions"
)

// Filter is a function that applies a certain processing technique within a
// defined segment of a image.
//
// Filters manipulate an image in-place, which avoid the need to create copies
// of the original image when dealing with multithread processing and merge
// then once the operations are finished.
//
// The image is the full in-memory image, and the bounds is the region of the
// image that the filter should operate on.
type Filter func(image *image.RGBA, bounds image.Rectangle)

// ConcurrentFilter defines the configuration for executing a filter
// concurrently by partitioning the image into segments and providing
// a mechanism to build per-partition filters.
type ConcurrentFilter struct {

	// Builder is a function that, given the full image and bounds for a segment,
	// returns a Filter adapted to process that segment. It can inject
	// stateful or context-specific behavior required for concurrent execution.
	Builder func(img *image.RGBA, bounds image.Rectangle) Filter

	// PartitionMethod defines how the image is divided into partitions for
	// concurrent processing. Different partitioning strategies can be used
	// to optimize for specific workloads.
	PartitionMethod partitions.ImagePartitioner
}

// FilterPipeline defines the entire processing flow of a filter, composed of
// an optional preprocessing step, the core filter and an optional closure
// function to adapt filters for concurrent workflows.
//
// A pipeline can recursively chain preprocessing steps, making it flexible for
// building complex operations by composing a series of simpler ones.
//
// When creating a new pipeline, please use the BuildPipelineFilter function,
// which ensures that the pipeline will be well formed.
type FilterPipeline struct {

	// Preprocess defines an optional preprocessing pipeline to be applied
	// once to the entire image before the main filter runs. It can itself be
	// another FilterPipeline, allowing for recursive chaining of operations.
	// If nil, no preprocessing operations are performed.
	Preprocess *FilterPipeline

	// Filter is the core filter applied in both serial (by default) and
	// concurrent execution modes. For concurrent execution, only the image
	// boundaries should be modified.
	Filter Filter

	// ConcurrentFilter defines the configuration needed for running the filter
	// concurrently, including the method for partitioning the image and
	// building per-partition filters.
	ConcurrentFilter ConcurrentFilter
}

// BuildFilterPipeline creates and initializes a FilterPipeline with the
// provided core filter, optional preprocessing step, concurrent builder
// function, and partitioning strategy.
//
// If partitioner is nil, the vertical partitioning method is used by default.
func BuildFilterPipeline(
	filter Filter,
	preprocess *FilterPipeline,
	concurrentBuilder func(img *image.RGBA, bounds image.Rectangle) Filter,
	partitioner partitions.ImagePartitioner) FilterPipeline {

	if partitioner == nil {
		partitioner = partitions.GetVerticalPartitions
	}

	return FilterPipeline{
		Preprocess: preprocess,
		Filter:     filter,
		ConcurrentFilter: ConcurrentFilter{
			Builder:         concurrentBuilder,
			PartitionMethod: partitioner,
		},
	}
}
