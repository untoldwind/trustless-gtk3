package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

// Format is a representation of Cairo's cairo_format_t.
type Format int

const (
	FormatInvalid   Format = C.CAIRO_FORMAT_INVALID
	FormatARGB32    Format = C.CAIRO_FORMAT_ARGB32
	FormatRGB24     Format = C.CAIRO_FORMAT_RGB24
	FormatA8        Format = C.CAIRO_FORMAT_A8
	FormatA1        Format = C.CAIRO_FORMAT_A1
	FormatRGB16_565 Format = C.CAIRO_FORMAT_RGB16_565
	FormatRGB30     Format = C.CAIRO_FORMAT_RGB30
)
