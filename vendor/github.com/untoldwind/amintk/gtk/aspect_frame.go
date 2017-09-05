package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import "unsafe"

// Frame is a representation of GTK's GtkFrame.
type AspectFrame struct {
	Frame
}

// native returns a pointer to the underlying GtkLevelBar.
func (v *AspectFrame) native() *C.GtkAspectFrame {
	if v == nil {
		return nil
	}
	return (*C.GtkAspectFrame)(v.Native())
}

// LayoutNew is a wrapper around gtk_layout_new().
func AspectFrameNew(label string, xAlign, yAlign, ratio float32, obeyChild bool) *AspectFrame {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	c := C.gtk_aspect_frame_new((*C.gchar)(cstr), C.gfloat(xAlign), C.gfloat(yAlign), C.gfloat(ratio), gbool(obeyChild))
	return wrapAspectFrame(unsafe.Pointer(c))
}

func wrapAspectFrame(p unsafe.Pointer) *AspectFrame {
	if frame := wrapFrame(p); frame != nil {
		return &AspectFrame{Frame: *frame}
	}
	return nil
}

func (v *AspectFrame) Set(xAlign, yAlign, ratio float32, obeyChild bool) {
	C.gtk_aspect_frame_set(v.native(), C.gfloat(xAlign), C.gfloat(yAlign), C.gfloat(ratio), gbool(obeyChild))
}
