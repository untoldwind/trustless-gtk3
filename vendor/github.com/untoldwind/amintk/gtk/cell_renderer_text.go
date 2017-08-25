package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import "unsafe"

// CellRendererText is a representation of GTK's GtkCellRendererText.
type CellRendererText struct {
	CellRenderer
}

// native returns a pointer to the underlying GtkTreeView.
func (v *CellRendererText) native() *C.GtkCellRendererText {
	if v == nil {
		return nil
	}
	return (*C.GtkCellRendererText)(v.Native())
}

// CellRendererTextNew is a wrapper around gtk_cell_renderer_text_new().
func CellRendererTextNew() *CellRendererText {
	c := C.gtk_cell_renderer_text_new()
	return wrapCellRendererText(unsafe.Pointer(c))
}

func wrapCellRendererText(p unsafe.Pointer) *CellRendererText {
	if cellRenderer := wrapCellRenderer(p); cellRenderer != nil {
		return &CellRendererText{CellRenderer: *cellRenderer}
	}
	return nil
}
