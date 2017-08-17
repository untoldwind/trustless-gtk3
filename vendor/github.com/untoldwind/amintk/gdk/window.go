package gdk

// #cgo pkg-config: gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Window is a representation of GDK's GdkWindow.
type Window struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkComboBox.
func (v *Window) native() *C.GdkWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GdkWindow)(p)
}

func (v *Window) SetCursor(cursor *Cursor) {
	C.gdk_window_set_cursor(v.native(), cursor.native())
}
