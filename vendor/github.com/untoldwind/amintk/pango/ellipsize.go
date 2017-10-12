package pango

// #cgo pkg-config: pango
// #include <pango/pango.h>
import "C"

// EllipsizeMode is a representation of Pango's PangoEllipsizeMode.
type EllipsizeMode int

const (
	EllipsizeModeNone   EllipsizeMode = C.PANGO_ELLIPSIZE_NONE
	EllipsizeModeStart  EllipsizeMode = C.PANGO_ELLIPSIZE_START
	EllipsizeModeMiddle EllipsizeMode = C.PANGO_ELLIPSIZE_MIDDLE
	EllipsizeModeEnd    EllipsizeMode = C.PANGO_ELLIPSIZE_END
)
