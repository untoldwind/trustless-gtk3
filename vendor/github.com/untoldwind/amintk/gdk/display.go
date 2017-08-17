package gdk

// #cgo pkg-config: gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Display is a representation of GDK's GdkDisplay.
type Display struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkComboBox.
func (v *Display) native() *C.GdkDisplay {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GdkDisplay)(p)
}

// DisplayGetDefault is a wrapper around gdk_display_get_default().
func DisplayGetDefault() *Display {
	c := C.gdk_display_get_default()
	obj := glib.WrapObject(unsafe.Pointer(c))
	return &Display{Object: obj}
}

func (v *Display) CursorFromName(name string) *Cursor {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gdk_cursor_new_from_name(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	obj := glib.WrapObject(unsafe.Pointer(c))
	return &Cursor{Object: obj}
}
