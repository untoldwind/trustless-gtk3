package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// TreePath is a representation of GTK's GtkTreePath.
type TreePath struct {
	GtkTreePath *C.GtkTreePath
}

// native returns a pointer to the underlying GtkTreePath.
func (v *TreePath) native() *C.GtkTreePath {
	if v == nil {
		return nil
	}
	return v.GtkTreePath
}

func (v *TreePath) free() {
	C.gtk_tree_path_free(v.native())
}

func wrapTreePath(c *C.GtkTreePath) *TreePath {
	if c == nil {
		return nil
	}
	t := &TreePath{c}
	runtime.SetFinalizer(t, (*TreePath).free)
	return t
}

// GetIndices is a wrapper around gtk_tree_path_get_indices_with_depth
func (v *TreePath) GetIndices() []int {
	var depth C.gint
	var goindices []int
	var ginthelp C.gint
	indices := uintptr(unsafe.Pointer(C.gtk_tree_path_get_indices_with_depth(v.native(), &depth)))
	size := unsafe.Sizeof(ginthelp)
	for i := 0; i < int(depth); i++ {
		goind := int(*((*C.gint)(unsafe.Pointer(indices))))
		goindices = append(goindices, goind)
		indices += size
	}
	return goindices
}

// String is a wrapper around gtk_tree_path_to_string().
func (v *TreePath) String() string {
	c := C.gtk_tree_path_to_string(v.native())
	return C.GoString((*C.char)(c))
}

// TreePathNewFromString is a wrapper around gtk_tree_path_new_from_string().
func TreePathNewFromString(path string) *TreePath {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_tree_path_new_from_string((*C.gchar)(cstr))
	return wrapTreePath(c)
}
