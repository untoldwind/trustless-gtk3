package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// InputPurpose is a representation of GTK's GtkInputPurpose.
type InputPurpose int

const (
	InputPurposeFreeForm InputPurpose = C.GTK_INPUT_PURPOSE_FREE_FORM
	InputPurposeAlpha    InputPurpose = C.GTK_INPUT_PURPOSE_ALPHA
	InputPurposeDigits   InputPurpose = C.GTK_INPUT_PURPOSE_DIGITS
	InputPurposeNumber   InputPurpose = C.GTK_INPUT_PURPOSE_NUMBER
	InputPurposePhone    InputPurpose = C.GTK_INPUT_PURPOSE_PHONE
	InputPurposeURL      InputPurpose = C.GTK_INPUT_PURPOSE_URL
	InputPurposeEmail    InputPurpose = C.GTK_INPUT_PURPOSE_EMAIL
	InputPurposeName     InputPurpose = C.GTK_INPUT_PURPOSE_NAME
	InputPurposePassword InputPurpose = C.GTK_INPUT_PURPOSE_PASSWORD
	InputPurposePIN      InputPurpose = C.GTK_INPUT_PURPOSE_PIN
)

// Entry is a representation of GTK's GtkEntry.
type Entry struct {
	Widget
}

// native returns a pointer to the underlying GtkEntry.
func (v *Entry) native() *C.GtkEntry {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkEntry)(p)
}

// EntryNew is a wrapper around gtk_entry_new().
func EntryNew() *Entry {
	c := C.gtk_entry_new()
	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapEntry(obj)
}

func wrapEntry(obj *glib.Object) *Entry {
	return &Entry{Widget{glib.InitiallyUnowned{Object: obj}}}
}

// SetInputPurpose is a wrapper around gtk_entry_set_input_purpose().
func (v *Entry) SetInputPurpose(purpose InputPurpose) {
	C.gtk_entry_set_input_purpose(v.native(), C.GtkInputPurpose(purpose))
}

// GetInputPurpose is a wrapper around gtk_entry_get_input_purpose().
func (v *Entry) GetInputPurpose() InputPurpose {
	c := C.gtk_entry_get_input_purpose(v.native())
	return InputPurpose(c)
}

// SetWidthChars is a wrapper around gtk_entry_set_width_chars().
func (v *Entry) SetWidthChars(nChars int) {
	C.gtk_entry_set_width_chars(v.native(), C.gint(nChars))
}

// SetText is a wrapper around gtk_entry_set_text().
func (v *Entry) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_entry_set_text(v.native(), (*C.gchar)(cstr))
}

// GetText is a wrapper around gtk_entry_get_text().
func (v *Entry) GetText() string {
	c := C.gtk_entry_get_text(v.native())
	return C.GoString((*C.char)(c))
}

// SetVisibility is a wrapper around gtk_entry_set_visibility().
func (v *Entry) SetVisibility(visible bool) {
	C.gtk_entry_set_visibility(v.native(), gbool(visible))
}

// GetVisibility is a wrapper around gtk_entry_get_visibility().
func (v *Entry) GetVisibility() bool {
	c := C.gtk_entry_get_visibility(v.native())
	return gobool(c)
}
