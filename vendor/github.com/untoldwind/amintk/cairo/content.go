package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

// Content is a representation of Cairo's cairo_content_t.
type Content int

const (
	ContentColor      Content = C.CAIRO_CONTENT_COLOR
	ContentAlpha      Content = C.CAIRO_CONTENT_ALPHA
	ContentColorAlpha Content = C.CAIRO_CONTENT_COLOR_ALPHA
)
