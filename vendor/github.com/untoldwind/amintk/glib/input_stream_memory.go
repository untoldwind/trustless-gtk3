package glib

// #cgo pkg-config: glib-2.0 gobject-2.0 gio-2.0
// #include "input_stream_memory.go.h"
import "C"
import "unsafe"

// MenuModel is a representation of GMenuModel.
type InputStreamMemory struct {
	InputStream
}

// native returns a pointer to the underlying GtkBox.
func (v *InputStreamMemory) native() *C.GMemoryInputStream {
	if v == nil {
		return nil
	}
	return (*C.GMemoryInputStream)(v.Native())
}

func wrapInputStreamMemory(p unsafe.Pointer) *InputStreamMemory {
	if inputStream := wrapInputStream(p); inputStream != nil {
		return &InputStreamMemory{InputStream: *inputStream}
	}
	return nil
}

func InputStreamMemoryNew() *InputStreamMemory {
	c := C.g_memory_input_stream_new()
	return wrapInputStreamMemory(unsafe.Pointer(c))
}

func InputStreamMemoryFromData(data []byte) *InputStreamMemory {
	cData := C.CString(string(data))
	c := C._g_memory_input_stream_new_from_data(unsafe.Pointer(cData), C.gssize(len(data)))
	return wrapInputStreamMemory(unsafe.Pointer(c))
}
