package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// StackSwitcher is a representation of GTK's GtkStackSwitcher
type StackSwitcher struct {
	Box
}

// native returns a pointer to the underlying GtkStackSwitcher.
func (v *StackSwitcher) native() *C.GtkStackSwitcher {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkStackSwitcher)(p)
}

// StackSwitcherNew is a wrapper around gtk_stack_switcher_new().
func StackSwitcherNew() *StackSwitcher {
	c := C.gtk_stack_switcher_new()
	return wrapStackSwitcher(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapStackSwitcher(obj *glib.Object) *StackSwitcher {
	return &StackSwitcher{Box{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}

// SetStack is a wrapper around gtk_stack_switcher_set_stack().
func (v *StackSwitcher) SetStack(stack *Stack) {
	C.gtk_stack_switcher_set_stack(v.native(), stack.native())
}

// GetStack is a wrapper around gtk_stack_switcher_get_stack().
func (v *StackSwitcher) GetStack() *Stack {
	c := C.gtk_stack_switcher_get_stack(v.native())
	return wrapStack(glib.WrapObject(unsafe.Pointer(c)))
}
