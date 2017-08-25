package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Container is a representation of GTK's GtkContainer.
type Container struct {
	Widget
}

// native returns a pointer to the underlying GtkContainer.
func (v *Container) native() *C.GtkContainer {
	if v == nil {
		return nil
	}
	return (*C.GtkContainer)(v.Native())
}

func wrapContainer(p unsafe.Pointer) *Container {
	if widget := wrapWidget(p); widget != nil {
		return &Container{Widget: *widget}
	}
	return nil
}

// Add is a wrapper around gtk_container_add().
func (v *Container) Add(w IWidget) {
	if w != nil {
		C.gtk_container_add(v.native(), w.toWidget())
	}
}

// Remove is a wrapper around gtk_container_remove().
func (v *Container) Remove(w IWidget) {
	if w != nil {
		C.gtk_container_remove(v.native(), w.toWidget())
	}
}

// GetFocusChild is a wrapper around gtk_container_get_focus_child().
func (v *Container) GetFocusChild() *Widget {
	c := C.gtk_container_get_focus_child(v.native())
	return wrapWidget(unsafe.Pointer(c))
}

// SetFocusChild is a wrapper around gtk_container_set_focus_child().
func (v *Container) SetFocusChild(child IWidget) {
	C.gtk_container_set_focus_child(v.native(), child.toWidget())
}

// GetBorderWidth is a wrapper around gtk_container_get_border_width().
func (v *Container) GetBorderWidth() uint {
	c := C.gtk_container_get_border_width(v.native())
	return uint(c)
}

// SetBorderWidth is a wrapper around gtk_container_set_border_width().
func (v *Container) SetBorderWidth(borderWidth uint) {
	C.gtk_container_set_border_width(v.native(), C.guint(borderWidth))
}
