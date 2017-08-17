package glib

// #cgo pkg-config: glib-2.0 gobject-2.0 gio-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// #include <gio/gio.h>
import "C"
import "unsafe"

// Menu is a representation of GMenu.
type Menu struct {
	MenuModel
}

// MenuNew is a wrapper around g_menu_new().
func MenuNew() *Menu {
	c := C.g_menu_new()
	return wrapMenu(WrapObject(unsafe.Pointer(c)))
}

func wrapMenu(obj *Object) *Menu {
	return &Menu{MenuModel{Object: obj}}
}
