package glib

// #cgo pkg-config: glib-2.0 gobject-2.0 gio-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// #include <gio/gio.h>
import "C"
import "unsafe"

// MenuModel is a representation of GMenuModel.
type InputStream struct {
	*Object
}

// native returns a pointer to the underlying GtkBox.
func (v *InputStream) native() *C.GInputStream {
	if v == nil {
		return nil
	}
	return (*C.GInputStream)(v.Native())
}

func wrapInputStream(p unsafe.Pointer) *InputStream {
	if obj := WrapObject(p); obj != nil {
		return &InputStream{Object: obj}
	}
	return nil
}
