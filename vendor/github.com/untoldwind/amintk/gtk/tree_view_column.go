package gtk

// #cgo pkg-config: gtk+-3.0
// #include "tree_view_column.go.h"
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// TreeViewColumns is a representation of GTK's GtkTreeViewColumn.
type TreeViewColumn struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkTreeViewColumn.
func (v *TreeViewColumn) native() *C.GtkTreeViewColumn {
	if v == nil {
		return nil
	}
	return (*C.GtkTreeViewColumn)(v.Native())
}

// TreeViewColumnNew is a wrapper around gtk_tree_view_column_new().
func TreeViewColumnNew() *TreeViewColumn {
	c := C.gtk_tree_view_column_new()
	return wrapTreeViewColumn(unsafe.Pointer(c))
}

// TreeViewColumnNewWithAttribute is a wrapper around
// gtk_tree_view_column_new_with_attributes() that only sets one
// attribute for one column.
func TreeViewColumnNewWithAttribute(title string, renderer ICellRenderer, attribute string, column int) *TreeViewColumn {
	t_cstr := C.CString(title)
	defer C.free(unsafe.Pointer(t_cstr))
	a_cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(a_cstr))
	c := C._gtk_tree_view_column_new_with_attributes_one((*C.gchar)(t_cstr),
		renderer.toCellRenderer(), (*C.gchar)(a_cstr), C.gint(column))
	return wrapTreeViewColumn(unsafe.Pointer(c))
}

func wrapTreeViewColumn(p unsafe.Pointer) *TreeViewColumn {
	if intiallyUnowned := glib.WrapInitiallyUnowned(p); intiallyUnowned != nil {
		return &TreeViewColumn{InitiallyUnowned: *intiallyUnowned}
	}
	return nil
}
