package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// TextBuffer is a representation of GTK's GtkTextBuffer.
type TextBuffer struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkTextBuffer.
func (v *TextBuffer) native() *C.GtkTextBuffer {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkTextBuffer)(p)
}

func wrapTextBuffer(obj *glib.Object) *TextBuffer {
	return &TextBuffer{Object: obj}
}

// GetStartIter() is a wrapper around gtk_text_buffer_get_start_iter().
func (v *TextBuffer) GetStartIter() *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_start_iter(v.native(), &iter)
	return (*TextIter)(&iter)
}

// GetEndIter() is a wrapper around gtk_text_buffer_get_end_iter().
func (v *TextBuffer) GetEndIter() *TextIter {
	var iter C.GtkTextIter
	C.gtk_text_buffer_get_end_iter(v.native(), &iter)
	return (*TextIter)(&iter)
}

func (v *TextBuffer) GetText(start, end *TextIter, includeHiddenChars bool) string {
	c := C.gtk_text_buffer_get_text(
		v.native(), (*C.GtkTextIter)(start), (*C.GtkTextIter)(end), gbool(includeHiddenChars),
	)
	gostr := C.GoString((*C.char)(c))
	C.g_free(C.gpointer(c))
	return gostr
}

func (v *TextBuffer) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_text_buffer_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}
