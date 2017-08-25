package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Switch is a representation of GTK's GtkSwitch.
type Switch struct {
	Widget
}

// native returns a pointer to the underlying GtkSwitch.
func (v *Switch) native() *C.GtkSwitch {
	if v == nil {
		return nil
	}
	return (*C.GtkSwitch)(v.Native())
}

// SwitchNew is a wrapper around gtk_switch_new().
func SwitchNew() *Switch {
	c := C.gtk_switch_new()
	return wrapSwitch(unsafe.Pointer(c))
}

func wrapSwitch(p unsafe.Pointer) *Switch {
	if widget := wrapWidget(p); widget != nil {
		return &Switch{Widget: *widget}
	}
	return nil
}

// GetActive is a wrapper around gtk_switch_get_active().
func (v *Switch) GetActive() bool {
	c := C.gtk_switch_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_switch_set_active().
func (v *Switch) SetActive(isActive bool) {
	C.gtk_switch_set_active(v.native(), gbool(isActive))
}
