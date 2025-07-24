package pipelines

import (
	"fmt"

	"dougdomingos.com/image-filters/filters"
)

type Step struct {
	Action *filters.Action
	Next *Step
}

type Recipe struct {
	head *Step
}

func (procQueue *Recipe) NextStep() *Step {
	if procQueue.head == nil {
		return nil
	}

	currentHead := procQueue.head
	procQueue.head = procQueue.head.Next
	return currentHead
}

func BuildRecipe(filterIDs []string) (Recipe, error) {
	var head, tail *Step

	for _, id := range filterIDs {
		pipeline, err := GetFilterPipeline(id)
		if err != nil {
			return Recipe{}, fmt.Errorf("[ERROR] No pipeline found for filter \"%q\"", id)
		}

		node := &Step{Action: &pipeline}

		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return Recipe{head: head}, nil
}
