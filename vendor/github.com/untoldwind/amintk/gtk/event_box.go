package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// EventBox is a representation of GTK's GtkEventBox.
type EventBox struct {
	Bin
}

// native returns a pointer to the underlying GtkEventBox.
func (v *EventBox) native() *C.GtkEventBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkEventBox)(p)
}

// EventBoxNew is a wrapper around gtk_event_box_new().
func EventBoxNew() *EventBox {
	c := C.gtk_event_box_new()
	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapEventBox(obj)
}

func wrapEventBox(obj *glib.Object) *EventBox {
	return &EventBox{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}
