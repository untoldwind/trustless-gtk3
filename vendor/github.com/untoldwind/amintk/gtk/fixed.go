package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import "unsafe"

// Layout is a representation of GTK's GtkLayout.
type Fixed struct {
	Container
}

// native returns a pointer to the underlying GtkLevelBar.
func (v *Fixed) native() *C.GtkFixed {
	if v == nil {
		return nil
	}
	return (*C.GtkFixed)(v.Native())
}

// LayoutNew is a wrapper around gtk_layout_new().
func FixedNew() *Fixed {
	c := C.gtk_fixed_new()
	return wrapFixed(unsafe.Pointer(c))
}

func wrapFixed(p unsafe.Pointer) *Fixed {
	if container := wrapContainer(p); container != nil {
		return &Fixed{Container: *container}
	}
	return nil
}

// Layout.Put is a wrapper around gtk_layout_put().
func (v *Fixed) Put(w IWidget, x, y int) {
	C.gtk_fixed_put(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// Layout.Move is a wrapper around gtk_layout_move().
func (v *Fixed) Move(w IWidget, x, y int) {
	C.gtk_fixed_move(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}
