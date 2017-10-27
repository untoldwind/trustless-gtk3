package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Button is a representation of GTK's GtkButton.
type Button struct {
	Bin
}

// native() returns a pointer to the underlying GtkButton.
func (v *Button) native() *C.GtkButton {
	if v == nil {
		return nil
	}
	return (*C.GtkButton)(v.Native())
}

// ButtonNew is a wrapper around gtk_button_new().
func ButtonNew() *Button {
	c := C.gtk_button_new()
	return wrapButton(unsafe.Pointer(c))
}

// ButtonNewWithLabel is a wrapper around gtk_button_new_with_label().
func ButtonNewWithLabel(label string) *Button {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_with_label((*C.gchar)(cstr))
	return wrapButton(unsafe.Pointer(c))
}

// ButtonNewFromIconName is a wrapper around gtk_button_new_from_icon_name().
func ButtonNewFromIconName(iconName string, size IconSize) *Button {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_button_new_from_icon_name((*C.gchar)(cstr),
		C.GtkIconSize(size))
	return wrapButton(unsafe.Pointer(c))
}

func wrapButton(p unsafe.Pointer) *Button {
	if bin := wrapBin(p); bin != nil {
		return &Button{Bin: *bin}
	}
	return nil
}

// SetLabel is a wrapper around gtk_button_set_label().
func (v *Button) SetLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_button_set_label(v.native(), (*C.gchar)(cstr))
}

// GetLabel is a wrapper around gtk_button_get_label().
func (v *Button) GetLabel() string {
	c := C.gtk_button_get_label(v.native())
	return C.GoString((*C.char)(c))
}

// SetAlwaysShowImage is a wrapper around gtk_button_set_always_show_image().
func (v *Button) SetAlwaysShowImage(alwaysShow bool) {
	C.gtk_button_set_always_show_image(v.native(), gbool(alwaysShow))
}

// GetAlwaysShowImage is a wrapper around gtk_button_get_always_show_image().
func (v *Button) GetAlwaysShowImage() bool {
	c := C.gtk_button_get_always_show_image(v.native())
	return gobool(c)
}

// SetImage is a wrapper around gtk_button_set_image().
func (v *Button) SetImage(image IWidget) {
	C.gtk_button_set_image(v.native(), image.toWidget())
}

// GetImage is a wrapper around gtk_button_get_image().
func (v *Button) GetImage() *Widget {
	c := C.gtk_button_get_image(v.native())
	return wrapWidget(unsafe.Pointer(c))
}

func (v *Button) OnClicked(callback func()) *glib.SignalHandle {
	return v.Connect("clicked", glib.CallbackVoidVoid(callback))
}
