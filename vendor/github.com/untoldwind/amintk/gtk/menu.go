package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/glib"
)

// Menu is a representation of GTK's GtkMenu.
type Menu struct {
	MenuShell
}

type IMenu interface {
	toMenu() *C.GtkMenu
	toWidget() *C.GtkWidget
}

// native() returns a pointer to the underlying GtkMenu.
func (v *Menu) native() *C.GtkMenu {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkMenu)(p)
}

func (v *Menu) toMenu() *C.GtkMenu {
	return v.native()
}

// MenuNew() is a wrapper around gtk_menu_new().
func MenuNew() *Menu {
	c := C.gtk_menu_new()
	return wrapMenu(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapMenu(obj *glib.Object) *Menu {
	return &Menu{MenuShell{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}

// PopupAtPointer is a wrapper for gtk_menu_popup_at_pointer(), on older versions it uses PopupAtMouseCursor
func (v *Menu) PopupAtPointer(triggerEvent *gdk.Event) {
	e := (*C.GdkEvent)(triggerEvent.Native())
	C.gtk_menu_popup_at_pointer(v.native(), e)
}
