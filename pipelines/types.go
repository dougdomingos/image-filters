package pipelines

import (
	"fmt"

	"dougdomingos.com/image-filters/filters/types"
)

// ProcessorNode represents a single filter pipeline in a processor queue. It
// may point to a subsequent pipeline in the queue, or nil if it's the last
// one.
type ProcessNode struct {

	// Pipeline holds the filter pipeline to be applied to the image.
	Pipeline *types.FilterPipeline

	// Next points to the next pipeline in the queue, or nil if there isn't
	// one.
	Next *ProcessNode
}

// ProcessorQueue represents a sequence of filter pipelines to be applied to a
// image.
type ProcessQueue struct {

	// head points to the first pipeline to be applied to the image.
	head *ProcessNode
}

// GetNode returns the next ProcessorNode to be executed by the filter engine.
// It does not preserves the current head of the queue, as ProcessQueues are
// discarded once all nodes are consumed.
func (procQueue *ProcessQueue) GetNode() *ProcessNode {
	if procQueue.head == nil {
		return nil
	}

	currentHead := procQueue.head
	procQueue.head = procQueue.head.Next
	return currentHead
}

// BuildProcessQueue creates a ProcessQueue containing one ProcessorNode for
// each requested filter. If a requested filter does not exist, the queue is
// discarded.
func BuildProcessQueue(filterIDs []string) (ProcessQueue, error) {
	var head, tail *ProcessNode

	for _, id := range filterIDs {
		pipeline, err := GetFilterPipeline(id)
		if err != nil {
			return ProcessQueue{}, fmt.Errorf("[ERROR] No pipeline found for filter \"%q\"", id)
		}

		node := &ProcessNode{Pipeline: &pipeline}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return ProcessQueue{head: head}, nil
}
