package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"

	"github.com/untoldwind/amintk/cairo"
)

// DrawingArea is a representation of GTK's GtkDrawingArea.
type DrawingArea struct {
	Widget
}

// native returns a pointer to the underlying GtkBox.
func (v *DrawingArea) native() *C.GtkDrawingArea {
	if v == nil {
		return nil
	}
	return (*C.GtkDrawingArea)(v.Native())
}

func wrapDrawingArea(p unsafe.Pointer) *DrawingArea {
	if widget := wrapWidget(p); widget != nil {
		return &DrawingArea{Widget: *widget}
	}
	return nil
}

// DrawingAreaNew is a wrapper around gtk_drawing_area_new().
func DrawingAreaNew() *DrawingArea {
	c := C.gtk_drawing_area_new()
	return wrapDrawingArea(unsafe.Pointer(c))
}

func (v *DrawingArea) OnDraw(callback func(context *cairo.Context) bool) *glib.SignalHandle {
	return v.Connect("draw", cairo.CallbackContextBoolean(callback))
}
