package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

type Popover struct {
	Bin
}

func (v *Popover) native() *C.GtkPopover {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkPopover)(p)
}

func PopoverNew(relative IWidget) *Popover {
	//Takes relative to widget
	var c *C.struct__GtkWidget
	if relative == nil {
		c = C.gtk_popover_new(nil)
	} else {
		c = C.gtk_popover_new(relative.toWidget())
	}
	return wrapPopover(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapPopover(obj *glib.Object) *Popover {
	return &Popover{Bin{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}}
}
