package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

// TextIter is a representation of GTK's GtkTextIter
type TextIter C.GtkTextIter

// native returns a pointer to the underlying GtkTextIter.
func (v *TextIter) native() *C.GtkTextIter {
	if v == nil {
		return nil
	}
	return (*C.GtkTextIter)(v)
}

// GetBuffer is a wrapper around gtk_text_iter_get_buffer().
func (v *TextIter) GetBuffer() *TextBuffer {
	c := C.gtk_text_iter_get_buffer(v.native())
	if c == nil {
		return nil
	}
	return wrapTextBuffer(unsafe.Pointer(c))
}

// GetText is a wrapper around gtk_text_iter_get_text().
func (v *TextIter) GetText(end *TextIter) string {
	c := C.gtk_text_iter_get_text(v.native(), end.native())
	return C.GoString((*C.char)(c))
}
