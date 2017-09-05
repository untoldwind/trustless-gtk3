package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

// FillRule is a representation of Cairo's cairo_fill_rule_t.
type FillRule int

const (
	FILL_RULE_WINDING  FillRule = C.CAIRO_FILL_RULE_WINDING
	FILL_RULE_EVEN_ODD FillRule = C.CAIRO_FILL_RULE_EVEN_ODD
)
