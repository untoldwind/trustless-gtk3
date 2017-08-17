package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// SelectionMode is a representation of GTK's GtkSelectionMode.
type SelectionMode int

const (
	SelectionModeNone     SelectionMode = C.GTK_SELECTION_NONE
	SelectionModeSingle   SelectionMode = C.GTK_SELECTION_SINGLE
	SelectionModeBrowse   SelectionMode = C.GTK_SELECTION_BROWSE
	SelectionModeMultiple SelectionMode = C.GTK_SELECTION_MULTIPLE
)

// ListBox is a representation of GTK's GtkListBox.
type ListBox struct {
	Container
}

// native returns a pointer to the underlying GtkListBox.
func (v *ListBox) native() *C.GtkListBox {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return (*C.GtkListBox)(p)
}

// ListBoxNew is a wrapper around gtk_list_box_new().
func ListBoxNew() *ListBox {
	c := C.gtk_list_box_new()
	return wrapListBox(glib.WrapObject(unsafe.Pointer(c)))
}

func wrapListBox(obj *glib.Object) *ListBox {
	return &ListBox{Container{Widget{glib.InitiallyUnowned{Object: obj}}}}
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
	if c == nil {
		return nil
	}
	return wrapListBoxRow(glib.WrapObject(unsafe.Pointer(c)))
}
