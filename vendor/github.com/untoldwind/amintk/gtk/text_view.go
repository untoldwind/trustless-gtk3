package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// TextView is a representation of GTK's GtkTextView
type TextView struct {
	Container
}

// native returns a pointer to the underlying GtkTextView.
func (v *TextView) native() *C.GtkTextView {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkTextView)(p)
}

// TextViewNew is a wrapper around gtk_text_view_new().
func TextViewNew() *TextView {
	c := C.gtk_text_view_new()
	return wrapTextView(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapTextView(obj *glib.Object) *TextView {
	return &TextView{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}
}

// GetBuffer is a wrapper around gtk_text_view_get_buffer().
func (v *TextView) GetBuffer() *TextBuffer {
	c := C.gtk_text_view_get_buffer(v.native())
	return wrapTextBuffer(glib.WrapObject(unsafe.Pointer(c)))
}

// SetBuffer is a wrapper around gtk_text_view_set_buffer().
func (v *TextView) SetBuffer(buffer *TextBuffer) {
	C.gtk_text_view_set_buffer(v.native(), buffer.native())
}
