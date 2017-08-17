package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Paned is a representation of GTK's GtkPaned.
type Paned struct {
	Bin
}

// native returns a pointer to the underlying GtkPaned.
func (v *Paned) native() *C.GtkPaned {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkPaned)(p)
}

// PanedNew is a wrapper around gtk_paned_new().
func PanedNew(orientation Orientation) *Paned {
	c := C.gtk_paned_new(C.GtkOrientation(orientation))
	return wrapPaned(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapPaned(obj *glib.Object) *Paned {
	return &Paned{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}

// Add1 is a wrapper around gtk_paned_add1().
func (v *Paned) Add1(child IWidget) {
	C.gtk_paned_add1(v.native(), child.toWidget())
}

// Add2 is a wrapper around gtk_paned_add2().
func (v *Paned) Add2(child IWidget) {
	C.gtk_paned_add2(v.native(), child.toWidget())
}

// SetPosition is a wrapper around gtk_paned_set_position().
func (v *Paned) SetPosition(position int) {
	C.gtk_paned_set_position(v.native(), C.gint(position))
}
