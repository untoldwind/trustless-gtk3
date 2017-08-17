package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// ComboBox is a representation of GTK's GtkComboBox.
type ComboBox struct {
	Bin
}

// native returns a pointer to the underlying GtkComboBox.
func (v *ComboBox) native() *C.GtkComboBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkComboBox)(p)
}

// ComboBoxNew is a wrapper around gtk_combo_box_new().
func ComboBoxNew() (*ComboBox, error) {
	c := C.gtk_combo_box_new()
	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapComboBox(obj), nil
}

func wrapComboBox(obj *glib.Object) *ComboBox {
	return &ComboBox{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}

// GetActive() is a wrapper around gtk_combo_box_get_active().
func (v *ComboBox) GetActive() int {
	c := C.gtk_combo_box_get_active(v.native())
	return int(c)
}

// SetActive() is a wrapper around gtk_combo_box_set_active().
func (v *ComboBox) SetActive(index int) {
	C.gtk_combo_box_set_active(v.native(), C.gint(index))
}
