package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Stack is a representation of GTK's GtkStack.
type Stack struct {
	Container
}

// native returns a pointer to the underlying GtkStack.
func (v *Stack) native() *C.GtkStack {
	if v == nil {
		return nil
	}
	return (*C.GtkStack)(v.Native())
}

// StackNew is a wrapper around gtk_stack_new().
func StackNew() *Stack {
	c := C.gtk_stack_new()
	return wrapStack(unsafe.Pointer(c))
}

func wrapStack(p unsafe.Pointer) *Stack {
	if container := wrapContainer(p); container != nil {
		return &Stack{Container: *container}
	}
	return nil
}

// AddNamed is a wrapper around gtk_stack_add_named().
func (v *Stack) AddNamed(child IWidget, name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_add_named(v.native(), child.toWidget(), (*C.gchar)(cstr))
}

// AddTitled is a wrapper around gtk_stack_add_titled().
func (v *Stack) AddTitled(child IWidget, name, title string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.gtk_stack_add_titled(v.native(), child.toWidget(), (*C.gchar)(cName),
		(*C.gchar)(cTitle))
}

// SetVisibleChild is a wrapper around gtk_stack_set_visible_child().
func (v *Stack) SetVisibleChild(child IWidget) {
	C.gtk_stack_set_visible_child(v.native(), child.toWidget())
}

// GetVisibleChild is a wrapper around gtk_stack_get_visible_child().
func (v *Stack) GetVisibleChild() *Widget {
	c := C.gtk_stack_get_visible_child(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(unsafe.Pointer(c))
}

// SetVisibleChildName is a wrapper around gtk_stack_set_visible_child_name().
func (v *Stack) SetVisibleChildName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_stack_set_visible_child_name(v.native(), (*C.gchar)(cstr))
}

// GetVisibleChildName is a wrapper around gtk_stack_get_visible_child_name().
func (v *Stack) GetVisibleChildName() string {
	c := C.gtk_stack_get_visible_child_name(v.native())
	return C.GoString((*C.char)(c))
}
