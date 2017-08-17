package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// Event is a representation of GDK's GdkEvent.
type Event struct {
	GdkEvent *C.GdkEvent
}

// native returns a pointer to the underlying GdkEvent.
func (v *Event) native() *C.GdkEvent {
	return v.GdkEvent
}

// Native returns a pointer to the underlying GdkEvent.
func (v *Event) Native() unsafe.Pointer {
	return unsafe.Pointer(v.native())
}

// WrapObject creates a new Object from a GObject pointer.
func WrapEvent(p unsafe.Pointer) *Event {
	if p == nil {
		return nil
	}

	event := &Event{GdkEvent: (*C.GdkEvent)(p)}
	runtime.SetFinalizer(event, finalizeEvent)

	return event
}

func finalizeEvent(event *Event) {
	C.gdk_event_free(event.native())
}
