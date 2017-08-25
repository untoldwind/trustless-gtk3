package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Paned is a representation of GTK's GtkPaned.
type Paned struct {
	Bin
}

// native returns a pointer to the underlying GtkPaned.
func (v *Paned) native() *C.GtkPaned {
	if v == nil {
		return nil
	}
	return (*C.GtkPaned)(v.Native())
}

// PanedNew is a wrapper around gtk_paned_new().
func PanedNew(orientation Orientation) *Paned {
	c := C.gtk_paned_new(C.GtkOrientation(orientation))
	return wrapPaned(unsafe.Pointer(c))
}

func wrapPaned(p unsafe.Pointer) *Paned {
	if bin := wrapBin(p); bin != nil {
		return &Paned{Bin: *bin}
	}
	return nil
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
