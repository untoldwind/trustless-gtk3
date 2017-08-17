package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Box is a representation of GTK's GtkBox.
type Box struct {
	Container
}

// native() returns a pointer to the underlying GtkBox.
func (v *Box) native() *C.GtkBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkBox)(p)
}

// BoxNew is a wrapper around gtk_box_new().
func BoxNew(orientation Orientation, spacing int) *Box {
	c := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapBox(obj)
}

func wrapBox(obj *glib.Object) *Box {
	return &Box{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}
}
