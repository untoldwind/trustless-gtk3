package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

// Antialias is a representation of Cairo's cairo_antialias_t.
type Antialias int

const (
	AntialiasDefault  Antialias = C.CAIRO_ANTIALIAS_DEFAULT
	AntialiasNone     Antialias = C.CAIRO_ANTIALIAS_NONE
	AntialiasGray     Antialias = C.CAIRO_ANTIALIAS_GRAY
	AntialiasSubpixel Antialias = C.CAIRO_ANTIALIAS_SUBPIXEL
	// ANTIALIAS_FAST     Antialias = C.CAIRO_ANTIALIAS_FAST (since 1.12)
	// ANTIALIAS_GOOD     Antialias = C.CAIRO_ANTIALIAS_GOOD (since 1.12)
	// ANTIALIAS_BEST     Antialias = C.CAIRO_ANTIALIAS_BEST (since 1.12)
)
