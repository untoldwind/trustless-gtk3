package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// EventBox is a representation of GTK's GtkEventBox.
type EventBox struct {
	Bin
}

// native returns a pointer to the underlying GtkEventBox.
func (v *EventBox) native() *C.GtkEventBox {
	if v == nil {
		return nil
	}
	return (*C.GtkEventBox)(v.Native())
}

// EventBoxNew is a wrapper around gtk_event_box_new().
func EventBoxNew() *EventBox {
	c := C.gtk_event_box_new()
	return wrapEventBox(unsafe.Pointer(c))
}

func wrapEventBox(p unsafe.Pointer) *EventBox {
	if bin := wrapBin(p); bin != nil {
		return &EventBox{Bin: *bin}
	}
	return nil
}
