package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Image is a representation of GTK's GtkImage.
type Image struct {
	Widget
}

// ImageNewFromIconName() is a wrapper around gtk_image_new_from_icon_name().
func ImageNewFromIconName(iconName string, size IconSize) *Image {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_image_new_from_icon_name((*C.gchar)(cstr),
		C.GtkIconSize(size))
	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapImage(obj)
}

func wrapImage(obj *glib.Object) *Image {
	return &Image{Widget{glib.InitiallyUnowned{Object: obj}}}
}
