package gtkextra

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtkextra.go.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
)

func ShowUri(screen *gdk.Screen, uri string) error {
	var cScreen *C.GdkScreen

	if screen != nil {
		cScreen = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	}
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))

	var err *C.GError = nil

	res := C.gtk_show_uri(cScreen, (*C.gchar)(cstr), C.gtk_get_current_event_time(), &err)
	if res == 0 {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}
