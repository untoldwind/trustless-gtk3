package gdk

// #cgo pkg-config: gdk-3.0
// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/untoldwind/amintk/glib"
)

// PixbufLoader is a representation of GDK's GdkPixbufLoader.
// Users of PixbufLoader are expected to call Close() when they are finished.
// PixbufLoader implements the io.WriteCloser interface
type PixbufLoader struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkBox.
func (v *PixbufLoader) native() *C.GdkPixbufLoader {
	if v == nil {
		return nil
	}
	return (*C.GdkPixbufLoader)(v.Native())
}

func WrapPixbufLoader(p unsafe.Pointer) *PixbufLoader {
	if obj := glib.WrapObject(p); obj != nil {
		return &PixbufLoader{Object: obj}
	}
	return nil
}

// PixbufLoaderNew is a wrapper around gdk_pixbuf_loader_new().
func PixbufLoaderNew() *PixbufLoader {
	c := C.gdk_pixbuf_loader_new()
	return WrapPixbufLoader(unsafe.Pointer(c))
}

// PixbufLoaderNewWithType is a wrapper around gdk_pixbuf_loader_new_with_type().
func PixbufLoaderNewWithType(t string) (*PixbufLoader, error) {
	var err *C.GError

	cstr := C.CString(t)
	defer C.free(unsafe.Pointer(cstr))

	c := C.gdk_pixbuf_loader_new_with_type((*C.char)(cstr), &err)
	if err != nil {
		defer C.g_error_free(err)
		return nil, errors.New(C.GoString((*C.char)(err.message)))
	}
	return WrapPixbufLoader(unsafe.Pointer(c)), nil
}

// Write is a wrapper around gdk_pixbuf_loader_write().  The
// function signature differs from the C equivalent to satisify the
// io.Writer interface.
func (v *PixbufLoader) Write(data []byte) (int, error) {
	// n is set to 0 on error, and set to len(data) otherwise.
	// This is a tiny hacky to satisfy io.Writer and io.WriteCloser,
	// which would allow access to all io and ioutil goodies,
	// and play along nice with go environment.

	if len(data) == 0 {
		return 0, nil
	}

	var err *C.GError
	c := C.gdk_pixbuf_loader_write(v.native(),
		(*C.guchar)(unsafe.Pointer(&data[0])), C.gsize(len(data)),
		&err)

	if !gobool(c) {
		defer C.g_error_free(err)
		return 0, errors.New(C.GoString((*C.char)(err.message)))
	}

	return len(data), nil
}

// Close is a wrapper around gdk_pixbuf_loader_close().  An error is
// returned instead of a bool like the native C function to support the
// io.Closer interface.
func (v *PixbufLoader) Close() error {
	var err *C.GError

	if ok := gobool(C.gdk_pixbuf_loader_close(v.native(), &err)); !ok {
		defer C.g_error_free(err)
		return errors.New(C.GoString((*C.char)(err.message)))
	}
	return nil
}

// GetPixbuf is a wrapper around gdk_pixbuf_loader_get_pixbuf().
func (v *PixbufLoader) GetPixbuf() *Pixbuf {
	c := C.gdk_pixbuf_loader_get_pixbuf(v.native())
	return WrapPixbuf(unsafe.Pointer(c))
}
