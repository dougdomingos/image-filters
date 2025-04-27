// Package partitions provides a set of image partitioning utility functions.
package partitions

import "image"

// ImagePartitioner is a function that splits a full image's boundaries into an
// arbitrary number of smaller partitions. These functions are used in
// concurrent executions to delegate different parts of a full image to
// different goroutines, depending on the process pipeline being used.
type ImagePartitioner func(bounds image.Rectangle, segments int) []image.Rectangle
