// Package imgutil provides a set of image utility functions.
package imgutil

import "image"

type PartitionCopier func(img *image.RGBA, bounds image.Rectangle, padding int) image.RGBA
