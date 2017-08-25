package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Overlay is a representation of GTK's GtkOverlay.
type Overlay struct {
	Bin
}

// native returns a pointer to the underlying GtkOverlay.
func (v *Overlay) native() *C.GtkOverlay {
	if v == nil {
		return nil
	}
	return (*C.GtkOverlay)(v.Native())
}

// OverlayNew is a wrapper around gtk_overlay_new().
func OverlayNew() *Overlay {
	c := C.gtk_overlay_new()
	return wrapOverlay(unsafe.Pointer(c))
}

func wrapOverlay(p unsafe.Pointer) *Overlay {
	if bin := wrapBin(p); bin != nil {
		return &Overlay{Bin: *bin}
	}
	return nil
}

// AddOverlay is a wrapper around gtk_overlay_add_overlay().
func (v *Overlay) AddOverlay(widget IWidget) {
	C.gtk_overlay_add_overlay(v.native(), widget.toWidget())
}
