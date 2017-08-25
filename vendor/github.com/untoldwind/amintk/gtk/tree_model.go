package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// TreeModelFlags is a representation of GTK's GtkTreeModelFlags.
type TreeModelFlags int

const (
	TreeModelFlagsItersPersist TreeModelFlags = C.GTK_TREE_MODEL_ITERS_PERSIST
	TreeModelFlagsListOnly     TreeModelFlags = C.GTK_TREE_MODEL_LIST_ONLY
)

// TreeModel is a representation of GTK's GtkTreeModel GInterface.
type TreeModel struct {
	*glib.Object
}

type ITreeModel interface {
	toTreeModel() *C.GtkTreeModel
}

// native returns a pointer to the underlying GObject as a GtkTreeModel.
func (v *TreeModel) native() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return (*C.GtkTreeModel)(v.Native())
}

func (v *TreeModel) toTreeModel() *C.GtkTreeModel {
	if v == nil {
		return nil
	}
	return v.native()
}

func wrapTreeModel(obj *glib.Object) *TreeModel {
	if obj != nil {
		return &TreeModel{Object: obj}
	}
	return nil
}

// GetFlags is a wrapper around gtk_tree_model_get_flags().
func (v *TreeModel) GetFlags() TreeModelFlags {
	c := C.gtk_tree_model_get_flags(v.native())
	return TreeModelFlags(c)
}

// GetNColumns is a wrapper around gtk_tree_model_get_n_columns().
func (v *TreeModel) GetNColumns() int {
	c := C.gtk_tree_model_get_n_columns(v.native())
	return int(c)
}

// GetColumnType is a wrapper around gtk_tree_model_get_column_type().
func (v *TreeModel) GetColumnType(index int) glib.Type {
	c := C.gtk_tree_model_get_column_type(v.native(), C.gint(index))
	return glib.Type(c)
}

// GetIter is a wrapper around gtk_tree_model_get_iter().
func (v *TreeModel) GetIter(path *TreePath) (*TreeIter, bool) {
	var iter C.GtkTreeIter
	c := C.gtk_tree_model_get_iter(v.native(), &iter, path.native())
	if !gobool(c) {
		return nil, false
	}
	return &TreeIter{iter}, true
}

// GetIterFromString is a wrapper around
// gtk_tree_model_get_iter_from_string().
func (v *TreeModel) GetIterFromString(path string) (*TreeIter, bool) {
	var iter C.GtkTreeIter
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_tree_model_get_iter_from_string(v.native(), &iter,
		(*C.gchar)(cstr))
	if !gobool(c) {
		return nil, false
	}
	return &TreeIter{iter}, true
}

// GetIterFirst is a wrapper around gtk_tree_model_get_iter_first().
func (v *TreeModel) GetIterFirst() (*TreeIter, bool) {
	var iter C.GtkTreeIter
	c := C.gtk_tree_model_get_iter_first(v.native(), &iter)
	if !gobool(c) {
		return nil, false
	}
	return &TreeIter{iter}, true
}

// GetPath is a wrapper around gtk_tree_model_get_path().
func (v *TreeModel) GetPath(iter *TreeIter) *TreePath {
	c := C.gtk_tree_model_get_path(v.native(), iter.native())
	return wrapTreePath(c)
}

// GetValue is a wrapper around gtk_tree_model_get_value().
func (v *TreeModel) GetValue(iter *TreeIter, column int) *glib.Value {
	val := glib.ValueAlloc()
	if val == nil {
		return nil
	}
	C.gtk_tree_model_get_value(
		(*C.GtkTreeModel)(unsafe.Pointer(v.native())),
		iter.native(),
		C.gint(column),
		(*C.GValue)(unsafe.Pointer(val.Native())))
	return val
}

// IterNext is a wrapper around gtk_tree_model_iter_next().
func (v *TreeModel) IterNext(iter *TreeIter) bool {
	c := C.gtk_tree_model_iter_next(v.native(), iter.native())
	return gobool(c)
}

// IterPrevious is a wrapper around gtk_tree_model_iter_previous().
func (v *TreeModel) IterPrevious(iter *TreeIter) bool {
	c := C.gtk_tree_model_iter_previous(v.native(), iter.native())
	return gobool(c)
}

// IterNthChild is a wrapper around gtk_tree_model_iter_nth_child().
func (v *TreeModel) IterNthChild(iter *TreeIter, parent *TreeIter, n int) bool {
	c := C.gtk_tree_model_iter_nth_child(v.native(), iter.native(), parent.native(), C.gint(n))
	return gobool(c)
}

// IterChildren is a wrapper around gtk_tree_model_iter_children().
func (v *TreeModel) IterChildren(iter, child *TreeIter) bool {
	var cIter, cChild *C.GtkTreeIter
	if iter != nil {
		cIter = iter.native()
	}
	cChild = child.native()
	c := C.gtk_tree_model_iter_children(v.native(), cChild, cIter)
	return gobool(c)
}

// IterNChildren is a wrapper around gtk_tree_model_iter_n_children().
func (v *TreeModel) IterNChildren(iter *TreeIter) int {
	var cIter *C.GtkTreeIter
	if iter != nil {
		cIter = iter.native()
	}
	c := C.gtk_tree_model_iter_n_children(v.native(), cIter)
	return int(c)
}
