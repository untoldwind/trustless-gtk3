package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

// LineCap is a representation of Cairo's cairo_line_cap_t.
type LineCap int

const (
	LineCapButt   LineCap = C.CAIRO_LINE_CAP_BUTT
	LineCapRound  LineCap = C.CAIRO_LINE_CAP_ROUND
	LineCapSquare LineCap = C.CAIRO_LINE_CAP_SQUARE
)
