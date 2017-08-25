package gtk

// #cgo pkg-config: gtk+-3.0
// #include "list_store.go.h"
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/glib"
)

// TreeStore is a representation of GTK's GtkTreeStore.
type TreeStore struct {
	*glib.Object

	// Interfaces
	TreeModel
}

// native returns a pointer to the underlying GtkListStore.
func (v *TreeStore) native() *C.GtkTreeStore {
	if v == nil {
		return nil
	}
	return (*C.GtkTreeStore)(v.Native())
}

func (v *TreeStore) toTreeModel() *C.GtkTreeModel {
	return (*C.GtkTreeModel)(v.Native())
}

// TreeStoreNew is a wrapper around gtk_tree_store_newv().
func TreeStoreNew(types ...glib.Type) *TreeStore {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_tree_store_newv(C.gint(len(types)), gtypes)
	return wrapTreeStore(unsafe.Pointer(c))
}

func wrapTreeStore(p unsafe.Pointer) *TreeStore {
	if obj := glib.WrapObject(p); obj != nil {
		tm := wrapTreeModel(obj)
		return &TreeStore{Object: obj, TreeModel: *tm}
	}
	return nil
}

// Append is a wrapper around gtk_tree_store_append().
func (v *TreeStore) Append(parent *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	var cParent *C.GtkTreeIter
	if parent != nil {
		cParent = parent.native()
	}
	C.gtk_tree_store_append(v.native(), &ti, cParent)
	iter := &TreeIter{ti}
	return iter
}

// Insert is a wrapper around gtk_tree_store_insert
func (v *TreeStore) Insert(parent *TreeIter, position int) *TreeIter {
	var ti C.GtkTreeIter
	var cParent *C.GtkTreeIter
	if parent != nil {
		cParent = parent.native()
	}
	C.gtk_tree_store_insert(v.native(), &ti, cParent, C.gint(position))
	iter := &TreeIter{ti}
	return iter
}

// SetValue is a wrapper around gtk_tree_store_set_value()
func (v *TreeStore) SetValue(iter *TreeIter, column int, value interface{}) error {
	switch value.(type) {
	case *gdk.Pixbuf:
		pix := value.(*gdk.Pixbuf)
		C._gtk_tree_store_set(v.native(), iter.native(), C.gint(column), unsafe.Pointer(pix.Native()))

	default:
		gv := glib.GValue(value)

		C.gtk_tree_store_set_value(v.native(), iter.native(),
			C.gint(column),
			(*C.GValue)(C.gpointer(gv.Native())))
	}
	return nil
}

// Remove is a wrapper around gtk_tree_store_remove().
func (v *TreeStore) Remove(iter *TreeIter) bool {
	var ti *C.GtkTreeIter
	if iter != nil {
		ti = iter.native()
	}
	return 0 != C.gtk_tree_store_remove(v.native(), ti)
}

// Clear is a wrapper around gtk_tree_store_clear().
func (v *TreeStore) Clear() {
	C.gtk_tree_store_clear(v.native())
}
