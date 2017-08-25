package gdk

// #cgo pkg-config: gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"
import (
	"github.com/untoldwind/amintk/glib"
)

// Cursor is a representation of GdkCursor.
type Cursor struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkComboBox.
func (v *Cursor) native() *C.GdkCursor {
	if v == nil {
		return nil
	}
	return (*C.GdkCursor)(v.Native())
}
