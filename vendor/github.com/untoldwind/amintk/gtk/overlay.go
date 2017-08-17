package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Overlay is a representation of GTK's GtkOverlay.
type Overlay struct {
	Bin
}

// native returns a pointer to the underlying GtkOverlay.
func (v *Overlay) native() *C.GtkOverlay {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkOverlay)(p)
}

// OverlayNew() is a wrapper around gtk_overlay_new().
func OverlayNew() *Overlay {
	c := C.gtk_overlay_new()
	return wrapOverlay(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapOverlay(obj *glib.Object) *Overlay {
	return &Overlay{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}

// AddOverlay() is a wrapper around gtk_overlay_add_overlay().
func (v *Overlay) AddOverlay(widget IWidget) {
	C.gtk_overlay_add_overlay(v.native(), widget.toWidget())
}
