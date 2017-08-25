package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
import "C"
import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

var typeEvent = glib.Type(C.gdk_event_get_type())

// Event is a representation of GDK's GdkEvent.
type Event struct {
	GdkEvent *C.GdkEvent
}

// native returns a pointer to the underlying GdkEvent.
func (v *Event) native() *C.GdkEvent {
	if v == nil {
		return nil
	}
	return v.GdkEvent
}

// Native returns a pointer to the underlying GdkEvent.
func (v *Event) Native() unsafe.Pointer {
	if v == nil {
		return nil
	}
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

type CallbackEventBoolean func(*Event) bool

func (c CallbackEventBoolean) Call(args []glib.Value) *glib.Value {
	var arg0 *Event
	var arg0Ok bool
	for _, value := range args {
		if actual, _ := value.Type(); actual == typeEvent {
			if p, ok := value.GetBoxed(); ok {
				arg0 = &Event{GdkEvent: (*C.GdkEvent)(p)}
				arg0Ok = true
			}
			break
		}
	}
	if !arg0Ok {
		fmt.Fprintln(os.Stderr, "WARNING: CallbackEventBoolean: No GdkEvent found in args")
	}
	res := c(arg0)
	return glib.GValue(res)
}
