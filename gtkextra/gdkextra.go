package gtkextra

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdkextra.go.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

func CursorNewFromName(display *gdk.Display, name string) (*gdk.Cursor, error) {
	cDisplay := (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gdk_cursor_new_from_name(cDisplay, (*C.gchar)(cstr))
	if c == nil {
		return nil, errors.New("gdk_cursor_new_from_name has nil result")
	}
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(c))}
	cursor := &gdk.Cursor{obj}
	obj.Ref()
	runtime.SetFinalizer(obj, (*glib.Object).Unref)

	return cursor, nil
}

func WindowSetCursor(window *gdk.Window, cursor *gdk.Cursor) {
	cWindow := (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	cCursor := (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	C.gdk_window_set_cursor(cWindow, cCursor)
}
