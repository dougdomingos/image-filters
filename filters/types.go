// Package filters provides a set of image filters implemented using
// pixel-based manipulation techniques.
package filters

import "image"

// Filter is a function that applies a specific effect to the provided image.
// All pixel operations are done in-place, and implementations are thread-safe
// by default.
type Filter func(image *image.RGBA)

type Action struct {
	Preprocess Filter
	Filter Filter
}
