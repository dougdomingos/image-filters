// Package filters provides a set of image filters implemented using
// pixel-based manipulation techniques.
package filters

import "image"

// Filter is a function that applies a certain processing technique within a
// defined segment of a image.
//
// Filters manipulate an image in-place, which avoid the need to create copies
// of the original image when dealing with multithread processing and merge
// then once the operations are finished.
//
// The `image` is the full in-memory image, and the `bounds` is the region of
// the image that the filter should operate on. The filter directly alters
// `img`.
type Filter func(image *image.RGBA, bounds image.Rectangle)
