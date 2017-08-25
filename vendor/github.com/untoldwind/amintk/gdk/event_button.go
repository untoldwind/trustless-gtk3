package gdk

// #cgo pkg-config: gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"
import "unsafe"

// EventButton is a representation of GDK's GdkEventButton.
type EventButton struct {
	*Event
}

func (v *EventButton) native() *C.GdkEventButton {
	if v == nil {
		return nil
	}
	return (*C.GdkEventButton)(unsafe.Pointer(v.Event.native()))
}

func (v *EventButton) Button() uint {
	c := v.native().button
	return uint(c)
}
