package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// ToggleButton is a representation of GTK's GtkToggleButton.
type ToggleButton struct {
	Button
}

// native returns a pointer to the underlying GtkToggleButton.
func (v *ToggleButton) native() *C.GtkToggleButton {
	if v == nil {
		return nil
	}
	return (*C.GtkToggleButton)(v.Native())
}

// ToggleButtonNew is a wrapper around gtk_toggle_button_new().
func ToggleButtonNew() *ToggleButton {
	c := C.gtk_toggle_button_new()
	return wrapToggleButton(unsafe.Pointer(c))
}

func wrapToggleButton(p unsafe.Pointer) *ToggleButton {
	if button := wrapButton(p); button != nil {
		return &ToggleButton{Button: *button}
	}
	return nil
}

// GetActive is a wrapper around gtk_toggle_button_get_active().
func (v *ToggleButton) GetActive() bool {
	c := C.gtk_toggle_button_get_active(v.native())
	return gobool(c)
}

// SetActive is a wrapper around gtk_toggle_button_set_active().
func (v *ToggleButton) SetActive(isActive bool) {
	C.gtk_toggle_button_set_active(v.native(), gbool(isActive))
}
