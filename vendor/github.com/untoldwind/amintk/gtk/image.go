package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/gdk"
)

// Image is a representation of GTK's GtkImage.
type Image struct {
	Widget
}

// ImageNew is a wrapper around gtk_image_new().
func ImageNew() *Image {
	c := C.gtk_image_new()
	return wrapImage(unsafe.Pointer(c))
}

func (v *Image) native() *C.GtkImage {
	if v == nil {
		return nil
	}
	return (*C.GtkImage)(v.Native())
}

// ImageNewFromIconName is a wrapper around gtk_image_new_from_icon_name().
func ImageNewFromIconName(iconName string, size IconSize) *Image {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_icon_name((*C.gchar)(cstr),
		C.GtkIconSize(size))
	return wrapImage(unsafe.Pointer(c))
}

// ImageNewFromPixbuf is a wrapper around gtk_image_new_from_pixbuf().
func ImageNewFromPixbuf(pixbuf *gdk.Pixbuf) *Image {
	ptr := (*C.GdkPixbuf)(pixbuf.Native())
	c := C.gtk_image_new_from_pixbuf(ptr)
	return wrapImage(unsafe.Pointer(c))
}

func wrapImage(p unsafe.Pointer) *Image {
	if widget := wrapWidget(p); widget != nil {
		return &Image{Widget: *widget}
	}
	return nil
}

// Clear() is a wrapper around gtk_image_clear().
func (v *Image) Clear() {
	C.gtk_image_clear(v.native())
}

// SetFromFixbuf is a wrapper around gtk_image_set_from_pixbuf().
func (v *Image) SetFromPixbuf(pixbuf *gdk.Pixbuf) {
	pbptr := (*C.GdkPixbuf)(pixbuf.Native())
	C.gtk_image_set_from_pixbuf(v.native(), pbptr)
}
