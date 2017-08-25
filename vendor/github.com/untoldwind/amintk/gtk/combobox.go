package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// ComboBox is a representation of GTK's GtkComboBox.
type ComboBox struct {
	Bin
}

// native returns a pointer to the underlying GtkComboBox.
func (v *ComboBox) native() *C.GtkComboBox {
	if v == nil {
		return nil
	}
	return (*C.GtkComboBox)(v.Native())
}

// ComboBoxNew is a wrapper around gtk_combo_box_new().
func ComboBoxNew() *ComboBox {
	c := C.gtk_combo_box_new()
	return wrapComboBox(unsafe.Pointer(c))
}

func wrapComboBox(p unsafe.Pointer) *ComboBox {
	if bin := wrapBin(p); bin != nil {
		return &ComboBox{Bin: *bin}
	}
	return nil
}

// GetActive is a wrapper around gtk_combo_box_get_active().
func (v *ComboBox) GetActive() int {
	c := C.gtk_combo_box_get_active(v.native())
	return int(c)
}

// SetActive is a wrapper around gtk_combo_box_set_active().
func (v *ComboBox) SetActive(index int) {
	C.gtk_combo_box_set_active(v.native(), C.gint(index))
}
