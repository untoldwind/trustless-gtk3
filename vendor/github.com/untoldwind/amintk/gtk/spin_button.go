package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// SpinButton is a representation of GTK's GtkSpinButton.
type SpinButton struct {
	Entry
}

// native returns a pointer to the underlying GtkSpinButton.
func (v *SpinButton) native() *C.GtkSpinButton {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkSpinButton)(p)
}

// SpinButtonNewWithRange() is a wrapper around
// gtk_spin_button_new_with_range().
func SpinButtonNewWithRange(min, max, step float64) *SpinButton {
	c := C.gtk_spin_button_new_with_range(C.gdouble(min), C.gdouble(max),
		C.gdouble(step))
	return wrapSpinButton(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapSpinButton(obj *glib.Object) *SpinButton {
	return &SpinButton{Entry{Widget{glib.InitiallyUnowned{Object: obj}}}}
}

// GetValueAsInt() is a wrapper around gtk_spin_button_get_value_as_int().
func (v *SpinButton) GetValueAsInt() int {
	c := C.gtk_spin_button_get_value_as_int(v.native())
	return int(c)
}

// SetValue() is a wrapper around gtk_spin_button_set_value().
func (v *SpinButton) SetValue(value float64) {
	C.gtk_spin_button_set_value(v.native(), C.gdouble(value))
}

// GetValue() is a wrapper around gtk_spin_button_get_value().
func (v *SpinButton) GetValue() float64 {
	c := C.gtk_spin_button_get_value(v.native())
	return float64(c)
}
