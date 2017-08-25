package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"github.com/untoldwind/amintk/glib"
)

// Orientation is a representation of GTK's GtkOrientation.
type Orientation int

const (
	OrientationHorizontal Orientation = C.GTK_ORIENTATION_HORIZONTAL
	OrientationVertical   Orientation = C.GTK_ORIENTATION_VERTICAL
)

// Orientable is a representation of GTK's GtkOrientable GInterface.
type Orientable struct {
	*glib.Object
}

// native returns a pointer to the underlying GObject as a GtkOrientable.
func (v *Orientable) native() *C.GtkOrientable {
	if v == nil {
		return nil
	}
	return (*C.GtkOrientable)(v.Native())
}

func wrapOrientable(obj *glib.Object) *Orientable {
	return &Orientable{Object: obj}
}

// GetOrientation() is a wrapper around gtk_orientable_get_orientation().
func (v *Orientable) GetOrientation() Orientation {
	c := C.gtk_orientable_get_orientation(v.native())
	return Orientation(c)
}

// SetOrientation() is a wrapper around gtk_orientable_set_orientation().
func (v *Orientable) SetOrientation(orientation Orientation) {
	C.gtk_orientable_set_orientation(v.native(),
		C.GtkOrientation(orientation))
}
