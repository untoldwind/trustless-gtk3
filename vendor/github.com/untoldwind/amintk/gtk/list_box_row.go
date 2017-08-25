package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"fmt"
	"os"
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

var typeListBoxRow = glib.Type(C.gtk_list_box_row_get_type())

// ListBoxRow is a representation of GTK's GtkListBoxRow.
type ListBoxRow struct {
	Bin
}

// native returns a pointer to the underlying GtkListBoxRow.
func (v *ListBoxRow) native() *C.GtkListBoxRow {
	if v == nil {
		return nil
	}
	return (*C.GtkListBoxRow)(v.Native())
}

func ListBoxRowNew() *ListBoxRow {
	c := C.gtk_list_box_row_new()
	return wrapListBoxRow(unsafe.Pointer(c))
}

func wrapListBoxRow(p unsafe.Pointer) *ListBoxRow {
	if bin := wrapBin(p); bin != nil {
		return &ListBoxRow{Bin: *bin}
	}
	return nil
}

// GetIndex is a wrapper around gtk_list_box_row_get_index()
func (v *ListBoxRow) GetIndex() int {
	c := C.gtk_list_box_row_get_index(v.native())
	return int(c)
}

type CallbackListBoxRowVoid func(*ListBoxRow)

func (c CallbackListBoxRowVoid) Call(args []glib.Value) *glib.Value {
	var arg0 *ListBoxRow
	var arg0Ok bool
	for _, value := range args {
		if actual, _ := value.Type(); actual == typeListBoxRow {
			if obj, ok := value.GetObject(); ok {
				arg0 = wrapListBoxRow(obj)
				arg0Ok = true
			}
			break
		}
	}
	if !arg0Ok {
		fmt.Fprintln(os.Stderr, "WARNING: CallbackListBoxRowVoid: No ListBoxRow found in args")
	}
	c(arg0)
	return nil
}
