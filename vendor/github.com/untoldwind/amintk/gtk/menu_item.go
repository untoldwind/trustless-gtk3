package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// MenuItem is a representation of GTK's GtkMenuItem.
type MenuItem struct {
	Bin
}

// IMenuItem is an interface type implemented by all structs
// embedding a MenuItem.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkMenuItem.
type IMenuItem interface {
	toMenuItem() *C.GtkMenuItem
	toWidget() *C.GtkWidget
}

// native returns a pointer to the underlying GtkMenuItem.
func (v *MenuItem) native() *C.GtkMenuItem {
	if v == nil {
		return nil
	}
	return (*C.GtkMenuItem)(v.Native())
}

func (v *MenuItem) toMenuItem() *C.GtkMenuItem {
	return v.native()
}

// MenuItemNew is a wrapper around gtk_menu_item_new().
func MenuItemNew() *MenuItem {
	c := C.gtk_menu_item_new()
	return wrapMenuItem(glib.WrapObject(unsafe.Pointer(c)))
}

// MenuItemNewWithLabel() is a wrapper around gtk_menu_item_new_with_label().
func MenuItemNewWithLabel(label string) *MenuItem {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_menu_item_new_with_label((*C.gchar)(cstr))
	return wrapMenuItem(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapMenuItem(obj *glib.Object) *MenuItem {
	return &MenuItem{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// Sets text on the menu_item label
func (v *MenuItem) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_menu_item_set_label(v.native(), (*C.gchar)(cstr))
}

// Gets text on the menu_item label
func (v *MenuItem) GetLabel() string {
	l := C.gtk_menu_item_get_label(v.native())
	return C.GoString((*C.char)(l))
}

func (v *MenuItem) OnActivate(callback func()) {
	if v != nil {
		v.Connect("activate", glib.CallbackVoidVoid(callback))
	}
}
