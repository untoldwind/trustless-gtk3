package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import "unsafe"

// Layout is a representation of GTK's GtkLayout.
type Layout struct {
	Container
}

// native returns a pointer to the underlying GtkLevelBar.
func (v *Layout) native() *C.GtkLayout {
	if v == nil {
		return nil
	}
	return (*C.GtkLayout)(v.Native())
}

// LayoutNew is a wrapper around gtk_layout_new().
func LayoutNew(hadjustment, vadjustment *Adjustment) *Layout {
	c := C.gtk_layout_new(hadjustment.native(), vadjustment.native())
	return wrapLayout(unsafe.Pointer(c))
}

func wrapLayout(p unsafe.Pointer) *Layout {
	if container := wrapContainer(p); container != nil {
		return &Layout{Container: *container}
	}
	return nil
}

// Layout.Put is a wrapper around gtk_layout_put().
func (v *Layout) Put(w IWidget, x, y int) {
	C.gtk_layout_put(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// Layout.Move is a wrapper around gtk_layout_move().
func (v *Layout) Move(w IWidget, x, y int) {
	C.gtk_layout_move(v.native(), w.toWidget(), C.gint(x), C.gint(y))
}

// Layout.SetSize is a wrapper around gtk_layout_set_size
func (v *Layout) SetSize(width, height uint) {
	C.gtk_layout_set_size(v.native(), C.guint(width), C.guint(height))
}

// Layout.GetSize is a wrapper around gtk_layout_get_size
func (v *Layout) GetSize() (width, height uint) {
	var w, h C.guint
	C.gtk_layout_get_size(v.native(), &w, &h)
	return uint(w), uint(h)
}
