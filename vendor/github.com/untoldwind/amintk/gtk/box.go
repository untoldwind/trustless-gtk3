package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Box is a representation of GTK's GtkBox.
type Box struct {
	Container
}

// native returns a pointer to the underlying GtkBox.
func (v *Box) native() *C.GtkBox {
	if v == nil {
		return nil
	}
	return (*C.GtkBox)(v.Native())
}

// BoxNew is a wrapper around gtk_box_new().
func BoxNew(orientation Orientation, spacing int) *Box {
	c := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	return wrapBox(unsafe.Pointer(c))
}

func wrapBox(p unsafe.Pointer) *Box {
	if container := wrapContainer(p); container != nil {
		return &Box{Container: *container}
	}
	return nil
}
