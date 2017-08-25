package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import "unsafe"

// MenuShell is a representation of GTK's GtkMenuShell.
type MenuShell struct {
	Container
}

// native returns a pointer to the underlying GtkMenuShell.
func (v *MenuShell) native() *C.GtkMenuShell {
	if v == nil {
		return nil
	}
	return (*C.GtkMenuShell)(v.Native())
}

func wrapMenuShell(p unsafe.Pointer) *MenuShell {
	if container := wrapContainer(p); container != nil {
		return &MenuShell{Container: *container}
	}
	return nil
}

// Append is a wrapper around gtk_menu_shell_append().
func (v *MenuShell) Append(child IMenuItem) {
	C.gtk_menu_shell_append(v.native(), child.toWidget())
}
