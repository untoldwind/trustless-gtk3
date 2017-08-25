package gtk

// #cgo pkg-config: gtk+-3.0
// #include "list_store.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/glib"
)

// ListStore is a representation of GTK's GtkListStore.
type ListStore struct {
	*glib.Object

	// Interfaces
	TreeModel
}

// native returns a pointer to the underlying GtkListStore.
func (v *ListStore) native() *C.GtkListStore {
	if v == nil {
		return nil
	}
	return (*C.GtkListStore)(v.Native())
}

// ListStoreNew is a wrapper around gtk_list_store_newv().
func ListStoreNew(types ...glib.Type) *ListStore {
	gtypes := C.alloc_types(C.int(len(types)))
	for n, val := range types {
		C.set_type(gtypes, C.int(n), C.GType(val))
	}
	defer C.g_free(C.gpointer(gtypes))
	c := C.gtk_list_store_newv(C.gint(len(types)), gtypes)

	return wrapListStore(unsafe.Pointer(c))
}

func (v *ListStore) toTreeModel() *C.GtkTreeModel {
	return (*C.GtkTreeModel)(v.Native())
}

func wrapListStore(p unsafe.Pointer) *ListStore {
	if obj := glib.WrapObject(p); obj != nil {
		tm := wrapTreeModel(obj)
		return &ListStore{Object: obj, TreeModel: *tm}
	}
	return nil
}

// Remove is a wrapper around gtk_list_store_remove().
func (v *ListStore) Remove(iter *TreeIter) bool {
	c := C.gtk_list_store_remove(v.native(), iter.native())
	return gobool(c)
}

// Set is a wrapper around gtk_list_store_set_value() but provides
// a function similar to gtk_list_store_set() in that multiple columns
// may be set by one call.  The length of columns and values slices must
// match, or Set() will return a non-nil error.
//
// As an example, a call to:
//  store.Set(iter, []int{0, 1}, []interface{}{"Foo", "Bar"})
// is functionally equivalent to calling the native C GTK function:
//  gtk_list_store_set(store, iter, 0, "Foo", 1, "Bar", -1);
func (v *ListStore) Set(iter *TreeIter, columns []int, values []interface{}) error {
	if len(columns) != len(values) {
		return errors.New("columns and values lengths do not match")
	}
	for i, val := range values {
		v.SetValue(iter, columns[i], val)
	}
	return nil
}

// SetValue is a wrapper around gtk_list_store_set_value().
func (v *ListStore) SetValue(iter *TreeIter, column int, value interface{}) error {
	switch value.(type) {
	case *gdk.Pixbuf:
		pix := value.(*gdk.Pixbuf)
		C._gtk_list_store_set(v.native(), iter.native(), C.gint(column), unsafe.Pointer(pix.Native()))

	default:
		gv := glib.GValue(value)

		C.gtk_list_store_set_value(v.native(), iter.native(),
			C.gint(column),
			(*C.GValue)(gv.Native()))
	}

	return nil
}

// InsertBefore is a wrapper around gtk_list_store_insert_before().
func (v *ListStore) InsertBefore(sibling *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_insert_before(v.native(), &ti, sibling.native())
	iter := &TreeIter{ti}
	return iter
}

// InsertAfter is a wrapper around gtk_list_store_insert_after().
func (v *ListStore) InsertAfter(sibling *TreeIter) *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_insert_after(v.native(), &ti, sibling.native())
	iter := &TreeIter{ti}
	return iter
}

// Prepend is a wrapper around gtk_list_store_prepend().
func (v *ListStore) Prepend() *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_prepend(v.native(), &ti)
	iter := &TreeIter{ti}
	return iter
}

// Append is a wrapper around gtk_list_store_append().
func (v *ListStore) Append() *TreeIter {
	var ti C.GtkTreeIter
	C.gtk_list_store_append(v.native(), &ti)
	iter := &TreeIter{ti}
	return iter
}

// Clear is a wrapper around gtk_list_store_clear().
func (v *ListStore) Clear() {
	C.gtk_list_store_clear(v.native())
}

// IterIsValid is a wrapper around gtk_list_store_iter_is_valid().
func (v *ListStore) IterIsValid(iter *TreeIter) bool {
	c := C.gtk_list_store_iter_is_valid(v.native(), iter.native())
	return gobool(c)
}

// Swap is a wrapper around gtk_list_store_swap().
func (v *ListStore) Swap(a, b *TreeIter) {
	C.gtk_list_store_swap(v.native(), a.native(), b.native())
}

// MoveBefore is a wrapper around gtk_list_store_move_before().
func (v *ListStore) MoveBefore(iter, position *TreeIter) {
	C.gtk_list_store_move_before(v.native(), iter.native(),
		position.native())
}

// MoveAfter is a wrapper around gtk_list_store_move_after().
func (v *ListStore) MoveAfter(iter, position *TreeIter) {
	C.gtk_list_store_move_after(v.native(), iter.native(),
		position.native())
}
