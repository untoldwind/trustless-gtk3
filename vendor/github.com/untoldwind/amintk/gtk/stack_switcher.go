package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// StackSwitcher is a representation of GTK's GtkStackSwitcher
type StackSwitcher struct {
	Box
}

// native returns a pointer to the underlying GtkStackSwitcher.
func (v *StackSwitcher) native() *C.GtkStackSwitcher {
	if v == nil {
		return nil
	}
	return (*C.GtkStackSwitcher)(v.Native())
}

// StackSwitcherNew is a wrapper around gtk_stack_switcher_new().
func StackSwitcherNew() *StackSwitcher {
	c := C.gtk_stack_switcher_new()
	return wrapStackSwitcher(unsafe.Pointer(c))
}

func wrapStackSwitcher(p unsafe.Pointer) *StackSwitcher {
	if box := wrapBox(p); box != nil {
		return &StackSwitcher{Box: *box}
	}
	return nil
}

// SetStack is a wrapper around gtk_stack_switcher_set_stack().
func (v *StackSwitcher) SetStack(stack *Stack) {
	C.gtk_stack_switcher_set_stack(v.native(), stack.native())
}

// GetStack is a wrapper around gtk_stack_switcher_get_stack().
func (v *StackSwitcher) GetStack() *Stack {
	c := C.gtk_stack_switcher_get_stack(v.native())
	return wrapStack(unsafe.Pointer(c))
}
