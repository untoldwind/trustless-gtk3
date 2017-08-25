package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// ComboBoxText is a representation of GTK's GtkComboBoxText.
type ComboBoxText struct {
	ComboBox
}

// native returns a pointer to the underlying GtkComboBoxText.
func (v *ComboBoxText) native() *C.GtkComboBoxText {
	if v == nil {
		return nil
	}
	return (*C.GtkComboBoxText)(v.Native())
}

// ComboBoxTextNew is a wrapper around gtk_combo_box_text_new().
func ComboBoxTextNew() *ComboBoxText {
	c := C.gtk_combo_box_text_new()
	return wrapComboBoxText(unsafe.Pointer(c))
}

func wrapComboBoxText(p unsafe.Pointer) *ComboBoxText {
	if comboBox := wrapComboBox(p); comboBox != nil {
		return &ComboBoxText{ComboBox: *comboBox}
	}
	return nil
}

// AppendText is a wrapper around gtk_combo_box_text_append_text().
func (v *ComboBoxText) AppendText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_append_text(v.native(), (*C.gchar)(cstr))
}

// PrependText is a wrapper around gtk_combo_box_text_prepend_text().
func (v *ComboBoxText) PrependText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_combo_box_text_prepend_text(v.native(), (*C.gchar)(cstr))
}

// RemoveAll is a wrapper around gtk_combo_box_text_remove_all().
func (v *ComboBoxText) RemoveAll() {
	C.gtk_combo_box_text_remove_all(v.native())
}
