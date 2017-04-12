package gtkextra

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtkextra.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gtk"
	"github.com/pkg/errors"
)

func ShowUriOnWindow(widget *gtk.Widget, uri string) error {
	toplevel, err := widget.GetToplevel()
	if err != nil {
		return errors.Wrap(err, "Failed to get toplevel")
	}
	if !toplevel.IsToplevel() {
		return errors.New("Toplevel is not toplevel (i.e. no window)")
	}

	cWindow := (*C.GtkWindow)(unsafe.Pointer(toplevel.Native()))
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))

	var gErr *C.GError

	res := C.gtk_show_uri_on_window(cWindow, cstr, C.gtk_get_current_event_time(), &gErr)
	if res == 0 {
		defer C.g_error_free(gErr)
		return errors.New(C.GoString((*C.char)(gErr.message)))
	}
	return nil
}
