package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// ToggleButton is a representation of GTK's GtkToggleButton.
type ToggleButton struct {
	Button
}

// native returns a pointer to the underlying GtkToggleButton.
func (v *ToggleButton) native() *C.GtkToggleButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkToggleButton)(p)
}

// ToggleButtonNew is a wrapper around gtk_toggle_button_new().
func ToggleButtonNew() *ToggleButton {
	c := C.gtk_toggle_button_new()
	return wrapToggleButton(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapToggleButton(obj *glib.Object) *ToggleButton {
	return &ToggleButton{Button{Bin{Container{Widget{
		glib.InitiallyUnowned{Object: obj}}}}}}
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
