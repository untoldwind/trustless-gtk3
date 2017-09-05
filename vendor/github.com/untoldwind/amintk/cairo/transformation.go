package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

// Translate is a wrapper around cairo_translate.
func (v *Context) Translate(tx, ty float64) {
	C.cairo_translate(v.native(), C.double(tx), C.double(ty))
}

// Scale is a wrapper around cairo_scale.
func (v *Context) Scale(sx, sy float64) {
	C.cairo_scale(v.native(), C.double(sx), C.double(sy))
}

// Rotate is a wrapper around cairo_rotate.
func (v *Context) Rotate(angle float64) {
	C.cairo_rotate(v.native(), C.double(angle))
}
