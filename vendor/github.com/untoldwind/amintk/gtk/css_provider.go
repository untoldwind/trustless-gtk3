package gtk

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// CssProvider is a representation of GTK's GtkCssProvider.
type CssProvider struct {
	*glib.Object
}

type IStyleProvider interface {
	toStyleProvider() *C.GtkStyleProvider
}

// native returns a pointer to the underlying GtkCssProvider.
func (v *CssProvider) native() *C.GtkCssProvider {
	if v == nil {
		return nil
	}
	return (*C.GtkCssProvider)(v.Native())
}

func (v *CssProvider) toStyleProvider() *C.GtkStyleProvider {
	if v == nil {
		return nil
	}
	return (*C.GtkStyleProvider)(unsafe.Pointer(v.native()))
}

// CssProviderNew is a wrapper around gtk_css_provider_new().
func CssProviderNew() *CssProvider {
	c := C.gtk_css_provider_new()
	return wrapCssProvider(unsafe.Pointer(c))
}

func wrapCssProvider(p unsafe.Pointer) *CssProvider {
	if obj := glib.WrapObject(p); obj != nil {
		return &CssProvider{Object: obj}
	}
	return nil
}

// LoadFromData is a wrapper around gtk_css_provider_load_from_data().
func (v *CssProvider) LoadFromData(data string) error {
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	var gerr *C.GError
	if C.gtk_css_provider_load_from_data(v.native(), (*C.gchar)(unsafe.Pointer(cdata)), C.gssize(len(data)), &gerr) == 0 {
		defer C.g_error_free(gerr)
		return errors.New(C.GoString((*C.char)(gerr.message)))
	}
	return nil
}
