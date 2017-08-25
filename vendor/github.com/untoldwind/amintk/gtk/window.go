package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"errors"
	"unsafe"
)

// WindowType is a representation of GTK's GtkWindowType.
type WindowType int

const (
	WindowToplevel WindowType = C.GTK_WINDOW_TOPLEVEL
	WindowPopup    WindowType = C.GTK_WINDOW_POPUP
)

// Window is a representation of GTK's GtkWindow.
type Window struct {
	Bin
}

// native returns a pointer to the underlying GtkWindow.
func (v *Window) native() *C.GtkWindow {
	if v == nil {
		return nil
	}
	return (*C.GtkWindow)(v.Native())
}

// WindowNew is a wrapper around gtk_window_new().
func WindowNew(t WindowType) *Window {
	c := C.gtk_window_new(C.GtkWindowType(t))
	return wrapWindow(unsafe.Pointer(c))
}

func wrapWindow(p unsafe.Pointer) *Window {
	if bin := wrapBin(p); bin != nil {
		return &Window{Bin: *bin}
	}
	return nil
}

// SetTitle is a wrapper around gtk_window_set_title().
func (v *Window) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_window_set_title(v.native(), (*C.gchar)(cstr))
}

// SetResizable is a wrapper around gtk_window_set_resizable().
func (v *Window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(v.native(), gbool(resizable))
}

// GetResizable is a wrapper around gtk_window_get_resizable().
func (v *Window) GetResizable() bool {
	c := C.gtk_window_get_resizable(v.native())
	return gobool(c)
}

// ActivateFocus is a wrapper around gtk_window_activate_focus().
func (v *Window) ActivateFocus() bool {
	c := C.gtk_window_activate_focus(v.native())
	return gobool(c)
}

// ActivateDefault is a wrapper around gtk_window_activate_default().
func (v *Window) ActivateDefault() bool {
	c := C.gtk_window_activate_default(v.native())
	return gobool(c)
}

// SetModal is a wrapper around gtk_window_set_modal().
func (v *Window) SetModal(modal bool) {
	C.gtk_window_set_modal(v.native(), gbool(modal))
}

// SetDefaultSize is a wrapper around gtk_window_set_default_size().
func (v *Window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(v.native(), C.gint(width), C.gint(height))
}

// GetDefaultSize is a wrapper around gtk_window_get_default_size().
func (v *Window) GetDefaultSize() (width, height int) {
	var w, h C.gint
	C.gtk_window_get_default_size(v.native(), &w, &h)
	return int(w), int(h)
}

func (v *Window) ShowUri(uri string) error {
	cstr := C.CString(uri)
	defer C.free(unsafe.Pointer(cstr))

	var gErr *C.GError

	res := C.gtk_show_uri_on_window(v.native(), cstr, C.gtk_get_current_event_time(), &gErr)
	if res == 0 {
		defer C.g_error_free(gErr)
		return errors.New(C.GoString((*C.char)(gErr.message)))
	}
	return nil
}
