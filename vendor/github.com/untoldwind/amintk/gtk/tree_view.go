package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// TreeView is a representation of GTK's GtkTreeView.
type TreeView struct {
	Container
}

// native returns a pointer to the underlying GtkTreeView.
func (v *TreeView) native() *C.GtkTreeView {
	if v == nil {
		return nil
	}
	return (*C.GtkTreeView)(v.Native())
}

// TreeViewNew is a wrapper around gtk_tree_view_new().
func TreeViewNew() *TreeView {
	c := C.gtk_tree_view_new()
	return wrapTreeView(unsafe.Pointer(c))
}

// TreeViewNewWithModel is a wrapper around gtk_tree_view_new_with_model().
func TreeViewNewWithModel(model ITreeModel) *TreeView {
	c := C.gtk_tree_view_new_with_model(model.toTreeModel())
	return wrapTreeView(unsafe.Pointer(c))
}

func wrapTreeView(p unsafe.Pointer) *TreeView {
	if container := wrapContainer(p); container != nil {
		return &TreeView{Container: *container}
	}
	return nil
}

// GetModel() is a wrapper around gtk_tree_view_get_model().
func (v *TreeView) GetModel() *TreeModel {
	c := C.gtk_tree_view_get_model(v.native())
	return wrapTreeModel(glib.WrapObject(unsafe.Pointer(c)))
}

// SetModel is a wrapper around gtk_tree_view_set_model().
func (v *TreeView) SetModel(model ITreeModel) {
	C.gtk_tree_view_set_model(v.native(), model.toTreeModel())
}

// GetSelection is a wrapper around gtk_tree_view_get_selection().
func (v *TreeView) GetSelection() *TreeSelection {
	c := C.gtk_tree_view_get_selection(v.native())
	return wrapTreeSelection(unsafe.Pointer(c))
}

// AppendColumn is a wrapper around gtk_tree_view_append_column().
func (v *TreeView) AppendColumn(column *TreeViewColumn) int {
	c := C.gtk_tree_view_append_column(v.native(), column.native())
	return int(c)
}
