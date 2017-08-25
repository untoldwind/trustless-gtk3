package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// Adjustment is a representation of GTK's GtkAdjustment.
type Adjustment struct {
	glib.InitiallyUnowned
}

// native returns a pointer to the underlying GtkAdjustment.
func (v *Adjustment) native() *C.GtkAdjustment {
	if v == nil {
		return nil
	}
	return (*C.GtkAdjustment)(v.Native())
}

// AdjustmentNew is a wrapper around gtk_adjustment_new().
func AdjustmentNew(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment {
	c := C.gtk_adjustment_new(C.gdouble(value),
		C.gdouble(lower),
		C.gdouble(upper),
		C.gdouble(stepIncrement),
		C.gdouble(pageIncrement),
		C.gdouble(pageSize))
	obj := glib.WrapObject(unsafe.Pointer(c))
	return wrapAdjustment(obj)
}

func wrapAdjustment(obj *glib.Object) *Adjustment {
	return &Adjustment{glib.InitiallyUnowned{Object: obj}}
}
