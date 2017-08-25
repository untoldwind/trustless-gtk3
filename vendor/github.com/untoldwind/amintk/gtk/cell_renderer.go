package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// CellRenderer is a representation of GTK's GtkCellRenderer.
type CellRenderer struct {
	glib.InitiallyUnowned
}

// ICellRenderer is an interface type implemented by all structs
// embedding a CellRenderer.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkCellRenderer.
type ICellRenderer interface {
	toCellRenderer() *C.GtkCellRenderer
}

func (v *CellRenderer) toCellRenderer() *C.GtkCellRenderer {
	if v == nil {
		return nil
	}
	return (*C.GtkCellRenderer)(v.Native())
}

func wrapCellRenderer(p unsafe.Pointer) *CellRenderer {
	if initiallyUnowned := glib.WrapInitiallyUnowned(p); initiallyUnowned != nil {
		return &CellRenderer{InitiallyUnowned: *initiallyUnowned}
	}
	return nil
}
