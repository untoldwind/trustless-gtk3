package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// ListBoxRow is a representation of GTK's GtkListBoxRow.
type ListBoxRow struct {
	Bin
}

// native returns a pointer to the underlying GtkListBoxRow.
func (v *ListBoxRow) native() *C.GtkListBoxRow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkListBoxRow)(p)
}

func ListBoxRowNew() *ListBoxRow {
	c := C.gtk_list_box_row_new()
	return wrapListBoxRow(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapListBoxRow(obj *glib.Object) *ListBoxRow {
	return &ListBoxRow{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}

// GetIndex is a wrapper around gtk_list_box_row_get_index()
func (v *ListBoxRow) GetIndex() int {
	c := C.gtk_list_box_row_get_index(v.native())
	return int(c)
}
