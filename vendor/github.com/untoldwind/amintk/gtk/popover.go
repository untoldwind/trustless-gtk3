package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
)

type Popover struct {
	Bin
}

func (v *Popover) native() *C.GtkPopover {
	if v == nil {
		return nil
	}
	return (*C.GtkPopover)(v.Native())
}

func PopoverNew(relative IWidget) *Popover {
	//Takes relative to widget
	var c *C.struct__GtkWidget
	if relative == nil {
		c = C.gtk_popover_new(nil)
	} else {
		c = C.gtk_popover_new(relative.toWidget())
	}
	return wrapPopover(unsafe.Pointer(c))
}

func wrapPopover(p unsafe.Pointer) *Popover {
	if bin := wrapBin(p); bin != nil {
		return &Popover{Bin: *bin}
	}
	return nil
}
