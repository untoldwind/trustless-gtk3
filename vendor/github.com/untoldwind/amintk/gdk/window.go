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
	if v == nil {
		return nil
	}
	return (*C.GdkWindow)(v.Native())
}

func WrapWindow(p unsafe.Pointer) *Window {
	if obj := glib.WrapObject(p); obj != nil {
		return &Window{Object: obj}
	}
	return nil
}

func (v *Window) SetCursor(cursor *Cursor) {
	C.gdk_window_set_cursor(v.native(), cursor.native())
}
