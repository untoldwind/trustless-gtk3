package gdk

// #cgo pkg-config: gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Cursor is a representation of GdkCursor.
type Cursor struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkComboBox.
func (v *Cursor) native() *C.GdkCursor {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GdkCursor)(p)
}
