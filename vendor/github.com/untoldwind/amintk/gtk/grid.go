package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Grid is a representation of GTK's GtkGrid.
type Grid struct {
	Container
	Orientable
}

// native returns a pointer to the underlying GtkGrid.
func (v *Grid) native() *C.GtkGrid {
	if v == nil {
		return nil
	}
	return (*C.GtkGrid)(v.Native())
}

// GridNew is a wrapper around gtk_grid_new().
func GridNew() *Grid {
	c := C.gtk_grid_new()
	return wrapGrid(unsafe.Pointer(c))
}

func wrapGrid(p unsafe.Pointer) *Grid {
	if container := wrapContainer(p); container != nil {
		o := wrapOrientable(container.Object)
		return &Grid{Container: *container, Orientable: *o}
	}
	return nil
}

// Attach is a wrapper around gtk_grid_attach().
func (v *Grid) Attach(child IWidget, left, top, width, height int) {
	C.gtk_grid_attach(v.native(), child.toWidget(), C.gint(left),
		C.gint(top), C.gint(width), C.gint(height))
}

// SetColumnSpacing is a wrapper around gtk_grid_set_column_spacing().
func (v *Grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(v.native(), C.guint(spacing))
}

// GetColumnSpacing is a wrapper around gtk_grid_get_column_spacing().
func (v *Grid) GetColumnSpacing() uint {
	c := C.gtk_grid_get_column_spacing(v.native())
	return uint(c)
}

// SetRowSpacing is a wrapper around gtk_grid_set_row_spacing().
func (v *Grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(v.native(), C.guint(spacing))
}

// GetRowSpacing is a wrapper around gtk_grid_get_row_spacing().
func (v *Grid) GetRowSpacing() uint {
	c := C.gtk_grid_get_row_spacing(v.native())
	return uint(c)
}
