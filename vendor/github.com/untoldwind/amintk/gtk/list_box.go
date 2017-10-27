package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// ListBox is a representation of GTK's GtkListBox.
type ListBox struct {
	Container
}

// native returns a pointer to the underlying GtkListBox.
func (v *ListBox) native() *C.GtkListBox {
	if v == nil {
		return nil
	}
	return (*C.GtkListBox)(v.Native())
}

// ListBoxNew is a wrapper around gtk_list_box_new().
func ListBoxNew() *ListBox {
	c := C.gtk_list_box_new()
	return wrapListBox(unsafe.Pointer(c))
}

func wrapListBox(p unsafe.Pointer) *ListBox {
	if container := wrapContainer(p); container != nil {
		return &ListBox{Container: *container}
	}
	return nil
}

// SetSelectionMode is a wrapper around gtk_list_box_set_selection_mode().
func (v *ListBox) SetSelectionMode(mode SelectionMode) {
	C.gtk_list_box_set_selection_mode(v.native(), C.GtkSelectionMode(mode))
}

// GetSelectionMode is a wrapper around gtk_list_box_get_selection_mode()
func (v *ListBox) GetSelectionMode() SelectionMode {
	c := C.gtk_list_box_get_selection_mode(v.native())
	return SelectionMode(c)
}

// SelectRow is a wrapper around gtk_list_box_select_row().
func (v *ListBox) SelectRow(row *ListBoxRow) {
	C.gtk_list_box_select_row(v.native(), row.native())
}

// GetSelectedRow is a wrapper around gtk_list_box_get_selected_row().
func (v *ListBox) GetSelectedRow() *ListBoxRow {
	c := C.gtk_list_box_get_selected_row(v.native())
	return wrapListBoxRow(unsafe.Pointer(c))
}

func (v *ListBox) OnAfterRowSelected(callback func(*ListBoxRow)) *glib.SignalHandle {
	return v.ConnectAfter("row-selected", CallbackListBoxRowVoid(callback))
}
