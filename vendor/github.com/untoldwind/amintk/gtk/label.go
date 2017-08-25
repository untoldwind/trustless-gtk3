package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// Label is a representation of GTK's GtkLabel.
type Label struct {
	Widget
}

// native returns a pointer to the underlying GtkLabel.
func (v *Label) native() *C.GtkLabel {
	if v == nil {
		return nil
	}
	return (*C.GtkLabel)(v.Native())
}

// LabelNew is a wrapper around gtk_label_new().
func LabelNew(str string) *Label {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_label_new((*C.gchar)(cstr))
	return wrapLabel(unsafe.Pointer(c))
}

func wrapLabel(p unsafe.Pointer) *Label {
	if widget := wrapWidget(p); widget != nil {
		return &Label{Widget: *widget}
	}
	return nil
}

// SetText is a wrapper around gtk_label_set_text().
func (v *Label) SetText(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_text(v.native(), (*C.gchar)(cstr))
}

// SetMarkup is a wrapper around gtk_label_set_markup().
func (v *Label) SetMarkup(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup(v.native(), (*C.gchar)(cstr))
}

// SetMarkupWithMnemonic is a wrapper around
// gtk_label_set_markup_with_mnemonic().
func (v *Label) SetMarkupWithMnemonic(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_markup_with_mnemonic(v.native(), (*C.gchar)(cstr))
}

// SetPattern is a wrapper around gtk_label_set_pattern().
func (v *Label) SetPattern(patern string) {
	cstr := C.CString(patern)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_label_set_pattern(v.native(), (*C.gchar)(cstr))
}

// SetSelectable is a wrapper around gtk_label_set_selectable().
func (v *Label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(v.native(), gbool(setting))
}
