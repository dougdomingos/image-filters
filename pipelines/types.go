package pipelines

import (
	"fmt"

	"dougdomingos.com/image-filters/filters"
	"dougdomingos.com/image-filters/filters/types"
)

// PipelineStep represents a single step in an image processing pipeline. Each
// step contains a filter and an reference to the next step, which may be nil
// if the current step is the last in the pipeline.
type PipelineStep struct {

	// Filter is the filter to be applied at this step.
	Filter *types.Filter

	// Next is a reference to the Step instance in the pipeline (nil if last)
	Next *PipelineStep
}

// Pipeline represents a sequence of filters to be applied in order. Filters
// are executed one-by-one through linked PipelineStep nodes. The pipeline
// behaves as a non-persistent queue, as it does not keep its steps as it's
// exhausted.
type Pipeline struct {

	// head is the first step in the pipeline.
	head *PipelineStep
}

// NextStep returns the current step in the pipeline and removes it from the
// pipeline, moving to the next step.
//
// If the pipeline has been fully consumed (i.e., Head is nil), it returns nil.
func (pipeline *Pipeline) NextStep() *PipelineStep {
	if pipeline.head == nil {
		return nil
	}

	currentHead := pipeline.head
	pipeline.head = pipeline.head.Next
	return currentHead
}

// IsEmpty checks if the current step at the start of the pipeline is nil.
func (pipeline *Pipeline) IsEmpty() bool {
	return pipeline.head == nil
}

// NewPipeline constructs a Pipeline from a list of filter IDs.
// Each ID is used to look up a corresponding filter via filters.GetFilter.
// If any ID is invalid or unrecognized, an error is returned.
//
// The resulting Pipeline can be used to sequentially apply each filter
// to an image in order.
func NewPipeline(filterIDs []string) (Pipeline, error) {
	var head, tail *PipelineStep

	for _, id := range filterIDs {
		pipeline, err := filters.GetFilter(id)
		if err != nil {
			return Pipeline{}, fmt.Errorf("[ERROR] No filter found for ID %q", id)
		}

		node := &PipelineStep{Filter: &pipeline}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return Pipeline{head: head}, nil
}
