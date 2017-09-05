package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import "unsafe"

// ShadowType is a representation of GTK's GtkShadowType.
type ShadowType int

const (
	ShadowTypeNone      ShadowType = C.GTK_SHADOW_NONE
	ShadowTypeIn        ShadowType = C.GTK_SHADOW_IN
	ShadowTypeOut       ShadowType = C.GTK_SHADOW_OUT
	ShadowTypeEtchedIn  ShadowType = C.GTK_SHADOW_ETCHED_IN
	ShadowTypeEtchedOut ShadowType = C.GTK_SHADOW_ETCHED_OUT
)

// Frame is a representation of GTK's GtkFrame.
type Frame struct {
	Bin
}

// native returns a pointer to the underlying GtkLevelBar.
func (v *Frame) native() *C.GtkFrame {
	if v == nil {
		return nil
	}
	return (*C.GtkFrame)(v.Native())
}

// LayoutNew is a wrapper around gtk_layout_new().
func FrameNew(label string) *Frame {
	var cstr *C.char
	if label != "" {
		cstr = C.CString(label)
		defer C.free(unsafe.Pointer(cstr))
	}
	c := C.gtk_frame_new((*C.gchar)(cstr))
	return wrapFrame(unsafe.Pointer(c))
}

func wrapFrame(p unsafe.Pointer) *Frame {
	if bin := wrapBin(p); bin != nil {
		return &Frame{Bin: *bin}
	}
	return nil
}

// SetLabel is a wrapper around gtk_frame_set_label().
func (v *Frame) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_frame_set_label(v.native(), (*C.gchar)(cstr))
}

// SetLabelWidget is a wrapper around gtk_frame_set_label_widget().
func (v *Frame) SetLabelWidget(labelWidget IWidget) {
	C.gtk_frame_set_label_widget(v.native(), labelWidget.toWidget())
}

// SetLabelAlign is a wrapper around gtk_frame_set_label_align().
func (v *Frame) SetLabelAlign(xAlign, yAlign float32) {
	C.gtk_frame_set_label_align(v.native(), C.gfloat(xAlign),
		C.gfloat(yAlign))
}

// SetShadowType is a wrapper around gtk_frame_set_shadow_type().
func (v *Frame) SetShadowType(t ShadowType) {
	C.gtk_frame_set_shadow_type(v.native(), C.GtkShadowType(t))
}

// GetLabel is a wrapper around gtk_frame_get_label().
func (v *Frame) GetLabel() string {
	c := C.gtk_frame_get_label(v.native())
	return C.GoString((*C.char)(c))
}

// GetLabelAlign is a wrapper around gtk_frame_get_label_align().
func (v *Frame) GetLabelAlign() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_frame_get_label_align(v.native(), &x, &y)
	return float32(x), float32(y)
}

// GetLabelWidget is a wrapper around gtk_frame_get_label_widget().
func (v *Frame) GetLabelWidget() *Widget {
	c := C.gtk_frame_get_label_widget(v.native())
	return wrapWidget(unsafe.Pointer(c))
}

// GetShadowType is a wrapper around gtk_frame_get_shadow_type().
func (v *Frame) GetShadowType() ShadowType {
	c := C.gtk_frame_get_shadow_type(v.native())
	return ShadowType(c)
}
