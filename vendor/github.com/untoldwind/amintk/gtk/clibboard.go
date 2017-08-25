package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/untoldwind/amintk/gdk"
	"github.com/untoldwind/amintk/glib"
)

// Clipboard is a wrapper around GTK's GtkClipboard.
type Clipboard struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkClipboard.
func (v *Clipboard) native() *C.GtkClipboard {
	if v == nil {
		return nil
	}
	return (*C.GtkClipboard)(v.Native())
}

// ClipboardGet is a wrapper around gtk_clipboard_get().
func ClipboardGet(atom gdk.Atom) *Clipboard {
	c := C.gtk_clipboard_get(C.GdkAtom(unsafe.Pointer(atom)))
	return wrapClipboard(unsafe.Pointer(c))
}

func wrapClipboard(p unsafe.Pointer) *Clipboard {
	if obj := glib.WrapObject(p); obj != nil {
		return &Clipboard{Object: obj}
	}
	return nil
}

// WaitIsTextAvailable is a wrapper around gtk_clipboard_wait_is_text_available
func (v *Clipboard) WaitIsTextAvailable() bool {
	c := C.gtk_clipboard_wait_is_text_available(v.native())
	return gobool(c)
}

// WaitForText is a wrapper around gtk_clipboard_wait_for_text
func (v *Clipboard) WaitForText() string {
	c := C.gtk_clipboard_wait_for_text(v.native())
	defer C.g_free(C.gpointer(c))
	return C.GoString((*C.char)(c))
}

// SetText() is a wrapper around gtk_clipboard_set_text().
func (v *Clipboard) SetText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_clipboard_set_text(v.native(), (*C.gchar)(cstr),
		C.gint(len(text)))
}
